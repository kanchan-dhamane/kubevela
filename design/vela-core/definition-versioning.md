# Versioning Support for KubeVela Definitions

<!-- toc -->
- [Versioning Support for KubeVela Definitions](#versioning-support-for-kubevela-definitions)
  - [Summary](#summary)
  - [Motivation](#motivation)
    - [Goals](#goals)
    - [Non-Goals](#non-goals)
  - [Acceptance Criteria](#acceptance-criteria)
  - [Current Implementation](#current-implementation)
      - [Reference:](#reference)
  - [Proposal 1](#proposal-1)
    - [Details](#details)
    - [Issues](#issues)
  - [Proposal 2](#proposal-2)
      - [ComponentDefinition CRD](#componentdefinition-crd)
      - [ComponentDefinition v2Alpha1 Object](#componentdefinition-v2alpha1-object)
    - [Implementation Details](#implementation-details)
      - [Reasons for selecting `v1Beta1` for storage:](#reasons-for-selecting-v1beta1-for-storage)
    - [Upgrading](#upgrading)
    - [How to handle existing components](#how-to-handle-existing-components)
    - [Notes/Considerations](#notesconsiderations)
    - [Summary of High Level Changes](#summary-of-high-level-changes)
      - [ComponentDefinition](#componentdefinition)
      - [DefintionRevision:](#defintionrevision)
      - [Application:](#application)
<!-- /toc -->

## Summary

Support explicit versioning for KubeVela Components and a way to specify which
version of the Component is to be used in an Application spec.


## Motivation

OAM/KubeVela Definitions (referred to as Definitions or Components for the rest
of the document) are the basic building blocks of the KubeVela platform. They
expose a contract similar to an API contract, which evolves from minor to major
versions. Definitions currently do have a `Revision` field in the Status, but it is an auto-incrementing integer.

Such a versioning scheme does not denote the type of the change (patch/bug, minor or major).

Applications are composed of Components that the KubeVela engine stitches
together. KubeVela automatically upgrades/reconciles the Application to the
latest when no Component Revision is specified. While we don't ideally want
Application developers to bother with such details, there are use cases where an automatic upgrade to the latest Component version is not desired.

- This can be resolved by specifying a Revision number when referring to the Component in an Application, but as the current versioning scheme does not provide hints on the type of change, it cannot be automated well.

### Goals

- Support Component versioning with Semantic Versions.
- Allow pinning, specific and non-specific versions of a Component in the
  Kubevela Application.

### Non-Goals

- Support for version range in Application. For eg. "type: my-component@>1.2.0"

## Acceptance Criteria

**User Story: Component version specification**

>**AS A** Component author\
>**I SHOULD** be able to publish every version of my Component with the Semantic Versioning scheme\
>**SO THAT** an Application developer can use a specific version of the Component.

**BDD Acceptance Criteria**

>**GIVEN** a Component spec\
>**AND** a version field of the spec set to V\
>**WHEN** the Component is applied to the KubeVela\
>**THEN** it should be listed as one of the many versions in the DefinitionRevision list

**User Story: Application Component version specification**

>**AS AN** Application developer\
>**I SHOULD** be able to specify a version (complete or partial) for every Component used\
>**SO THAT** I can control which version are deployed.

**BDD Acceptance Criteria**

>Scenario: Use the version specified in the Application manifest when deploying the service\
>**GIVEN** a Component A with version 1.2.3\
>**AND** a Component B with version 4.5.6\
>**AND** an Application composed of A 1.2.2 and B 4.4.2\
>**WHEN** the Application is deployed\
>**THEN** it uses Component A 1.2.2 and B 4.4.2

>**Variant:** Use the latest version for the part of the semVer that is not specified.\
>**AND** an Application composed of A 1.2 and B 4\
>**WHEN** the Application is deployed\
>**THEN** it uses Component A 1.2.3 and B 4.5.6

## Current Implementation

Currently, Kubevela has some support for controlling Definition versions based on K8s annotations and DefinitionRevisions. The annotation `definitionrevision.oam.dev/name` can be used to version the ComponentDefinition. For example if the following annotation is added to a ComponentDefinition, it produces a new DefinitionRevision and names the ComponentDefinition as `component-name-v4.4` .

> definitionrevision.oam.dev/name: "4.4"

![version](./kubevela-version.png)

This Component can then be referred in the Application as either of the following:

>"component-name@v4.4" - `NamedDefinitionRevision`\
>"component-name@v3"   - `DefinitionRevision`

This versioning scheme, although convenient, has the following issues:

- Since, there are more than one ways to reference a Component (`NamedDefinitionRevision` or `DefinitionRevision`), it proves difficult to maintain version consistency across clusters.
- Each change in the ComponentDefintion Spec creates a new DefinitionRevision only if the annotation `definitionrevision.oam.dev/name` is not present at all. If there is change in the ComponentDefintion Spec when the annotation  `definitionrevision.oam.dev/name` is present and not updated, the DefinitionRevision is updated in place. This behavior adds to the inconsistency.

#### Reference: 

https://kubevela.io/docs/platform-engineers/x-def-version/


## Proposal 1

### Details

This implementation, suggested [here](https://github.com/kubevela/kubevela/issues/6435#issuecomment-1892372596), proposes adding another explicit annotation for keeping track of ComponentDefinition Revisions. For example, we can add an annotation like `definitionrevision.oam.dev/version` and use that to generate the ComponentDefinition Revisions.

### Issues

The following issues assume adherence to strict backward compatibility.
- It might be very confusing for the users, because there is no easy way to differentiate or recognize how the Revision was generated.
- If both the annotations  `definitionrevision.oam.dev/name` and `definitionrevision.oam.dev/version` are present, there is no clear way to decide the Revision behavior.


## Proposal 2

This proposal targets the following major changes:

- Add a new API version(`v2Alpha1`) for ComponentDefinition CRD to support the new versioning model. Current Components should continue to work as is.
- Use Semantic versioning to identify new versions of a ComponentDefinition.
- Use the [Kubernetes API versioning model](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definition-versioning/) for managing the ComponentDefinition versions, which is to store an array of all version specifications. Each individual element in the array `spec.versions`  will also include a new field called `version` to indicate a specific Component version.
- Kubevela Application users can refer to a particular version of a ComponentDefinition like `component-name@component-version`. The `version` must be a valid Semantic version without the prefix `v` for example, `1.1.0`, `1.2.0-rc.0`.

#### ComponentDefinition CRD
![comp-def-crd](./comp-def-crd.png)

#### ComponentDefinition v2Alpha1 Object
![comp-def-instance](./comp-def-instance.png)


The Application will then have the ability to refer to the version when using a
Definition like this
```
    apiVersion: core.oam.dev/v1beta1
    kind: Application
    metadata:
        name: app-with-comp-versioning
    spec:
        components:
            - name: backend
              type: cloud-native-postgres-beta2@v3
```

### Implementation Details

Kubernetes only supports storing one API Version per CRD. We are targeting the current( `v1Beta1`) API version of `ComponentDefinition` for storage. We will need to implement a [Conversion webhook](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definition-versioning/#webhook-conversion) to translate the incoming objects of `v2Alpha1` to `v1Beta1`. It essentially means that even when we add `ComponentDefinitions` of the new API Version(`v2Alpha1`), only the translated object to `v1Beta1` ever gets stored in the database.

#### Reasons for selecting `v1Beta1` for storage:

-  We don't need to migrate existing `ComponentDefinition`'s to the new API Version(`v2Alpha1`) as part of the Kubevela upgrade.
-  In the interest of backward compatibility and stability, it is generally [advisable](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api_changes.md#operational-overview) to mark the most stable API Version of a CRD for storage. [Crossplane](https://github.com/crossplane-contrib/provider-upjet-aws/blob/main/apis/ec2/v1beta1/zz_instance_types.go#L1348 ) support for multiple versions also implements this.


### Upgrading

When selecting a Component version in the Application spec, it will be possible to
select a non-specific version, like `v1.0` or `v1`. KubeVela will
auto-upgrade the application to use the latest version in the specified version
series. For instance, if the application specifies `type: my-component-type@1.0`
and `v1.0.1` of the Definition is available, KubeVela will re-render the
Application using this version.

In case no Component version is specified, the latest exact version available at the time of Application creation will be pinned.

KubeVela will initiate the reconciliation of the Application as soon as a new
version of a Component is available and the Application is eligible (based on
the version specificity) to be upgraded.

> **Proposal Notes:** We are intentionally not planning to add a flag like `auto-upgrade:
> true|false` if one does not want to auto-upgrade their component. Application developers 
> should always use specific Semantic versions in the Application spec to disable auto upgrades. 
> This is also consistent
> with existing behavior where we always use the latest version of the
> Definition when the Application reconciles. We are just providing a way to opt
> out of this behavior by pinning the component version in the Application.
> This is also in line with the philosophy of keeping complexity at the platform
> level instead of the Application level.

> **For Definition Maintainers:** Ideally the upgrades to Definitions
> should be backward compatible for a given Major version. Updates to
> Definitions should never force the Application spec to change. If a 
> Component is changing something significant, it should be a new Definition.

### How to handle existing components

- We will treat the existing `v1Beta1` ComponentDefinition API Version as the legacy API version. These Components will continue to behave the same way, including their usage in an Application. This also applies to Applications & Components using the current `NamedDefinitionRevision` or `DefinitionRevision` versioning.
- An upgrade to KubeVela with this new versioning scheme support will automatically make it available for use. There will be no need for an explicit declaration.
- Conversion of Components between ComponentDefinition API Versions will be blocked. For example, a Component defined using the `v1Beta1` API cannot be edited to `v2Alpha1` API version or vice-versa, even though due to K8s behavior all objects will be displayed after translation to the latest API version. This translation will be implicit and will be handled by the Conversion webhooks.
  > This is still under consideration and might not be enforced depending on the implementation difficulty/limitations.
- All auto upgrade behavior will only apply to `v2Alpha1` API Components and will **not** need to be enabled via deployment flags.


### Notes/Considerations
- K8s GET resource call always returns the latest version by default ( https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definition-versioning/#version-priority ). When `v1Beta1` objects are accessed, the conversion webhook will be used to translate the result to `v2Alpha1` before being returned.
- When a new Component is created with `v2Alpha1` API target, the following objects will be stored:
  - A Component with `v1Beta1` API target which only includes the Spec of the highest version from the `v2Alpha1` object.
  - A DefinitionRevision for each Component version in the `v2Alpha1`  object will be created/updated.
- We need to be able to store all  the information of the new API Version object. One way to do this is via adding a [new annotation](https://github.com/shivi28/kube-vaccine/blob/main/api/v2/conversion.go).
- Conversion webhooks work via a [Hub and Spoke model](https://book.kubebuilder.io/multiversion-tutorial/conversion-concepts). It is not required that the API Version stored in the database and the API version marked as the Hub should be the same.The Hub can then be marked as per ease of implementation.


### Summary of High Level Changes

#### ComponentDefinition
- Add a new CRD version.
  - `spec.versions` is an array of all available/supported Component versions.
  - `status` for the ComponentDefintion will be modified to store the version metadata.
  
#### DefintionRevision:
- We are planning to keep the syntax for referring to a version of a Component in an Application.
- `v2Alpha1` Components will only be referable by the `DefinitionRevision` version and not the Revision. The Component version will continue to be appended to the `DefinitionRevision` Name. 

#### Application:
- Update the Component version parsing for `v2Alpha1` Components, to allow for auto upgrades.
