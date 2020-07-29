/**
 * Kore API
 * Kore API provides the frontend API (kore.appvia.io)
 *
 * The version of the OpenAPI document: 0.0.1
 * Contact: info@appvia.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1NamespaceClaimSpec from './V1NamespaceClaimSpec';
import V1NamespaceClaimStatus from './V1NamespaceClaimStatus';
import V1ObjectMeta from './V1ObjectMeta';

/**
 * The V1NamespaceClaim model module.
 * @module model/V1NamespaceClaim
 * @version 0.0.1
 */
class V1NamespaceClaim {
    /**
     * Constructs a new <code>V1NamespaceClaim</code>.
     * @alias module:model/V1NamespaceClaim
     */
    constructor() { 
        
        V1NamespaceClaim.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1NamespaceClaim</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1NamespaceClaim} obj Optional instance to populate.
     * @return {module:model/V1NamespaceClaim} The populated <code>V1NamespaceClaim</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1NamespaceClaim();

            if (data.hasOwnProperty('apiVersion')) {
                obj['apiVersion'] = ApiClient.convertToType(data['apiVersion'], 'String');
            }
            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('metadata')) {
                obj['metadata'] = V1ObjectMeta.constructFromObject(data['metadata']);
            }
            if (data.hasOwnProperty('spec')) {
                obj['spec'] = V1NamespaceClaimSpec.constructFromObject(data['spec']);
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = V1NamespaceClaimStatus.constructFromObject(data['status']);
            }
        }
        return obj;
    }

/**
     * Returns APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * @return {String}
     */
    getApiVersion() {
        return this.apiVersion;
    }

    /**
     * Sets APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * @param {String} apiVersion APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    setApiVersion(apiVersion) {
        this['apiVersion'] = apiVersion;
    }
/**
     * Returns Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * @return {String}
     */
    getKind() {
        return this.kind;
    }

    /**
     * Sets Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * @param {String} kind Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    setKind(kind) {
        this['kind'] = kind;
    }
/**
     * @return {module:model/V1ObjectMeta}
     */
    getMetadata() {
        return this.metadata;
    }

    /**
     * @param {module:model/V1ObjectMeta} metadata
     */
    setMetadata(metadata) {
        this['metadata'] = metadata;
    }
/**
     * @return {module:model/V1NamespaceClaimSpec}
     */
    getSpec() {
        return this.spec;
    }

    /**
     * @param {module:model/V1NamespaceClaimSpec} spec
     */
    setSpec(spec) {
        this['spec'] = spec;
    }
/**
     * @return {module:model/V1NamespaceClaimStatus}
     */
    getStatus() {
        return this.status;
    }

    /**
     * @param {module:model/V1NamespaceClaimStatus} status
     */
    setStatus(status) {
        this['status'] = status;
    }

}

/**
 * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
 * @member {String} apiVersion
 */
V1NamespaceClaim.prototype['apiVersion'] = undefined;

/**
 * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
 * @member {String} kind
 */
V1NamespaceClaim.prototype['kind'] = undefined;

/**
 * @member {module:model/V1ObjectMeta} metadata
 */
V1NamespaceClaim.prototype['metadata'] = undefined;

/**
 * @member {module:model/V1NamespaceClaimSpec} spec
 */
V1NamespaceClaim.prototype['spec'] = undefined;

/**
 * @member {module:model/V1NamespaceClaimStatus} status
 */
V1NamespaceClaim.prototype['status'] = undefined;






export default V1NamespaceClaim;

