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
import V1ObjectMeta from './V1ObjectMeta';
import V1SecurityOverviewSpec from './V1SecurityOverviewSpec';

/**
 * The V1SecurityOverview model module.
 * @module model/V1SecurityOverview
 * @version 0.0.1
 */
class V1SecurityOverview {
    /**
     * Constructs a new <code>V1SecurityOverview</code>.
     * @alias module:model/V1SecurityOverview
     */
    constructor() { 
        
        V1SecurityOverview.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1SecurityOverview</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1SecurityOverview} obj Optional instance to populate.
     * @return {module:model/V1SecurityOverview} The populated <code>V1SecurityOverview</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1SecurityOverview();

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
                obj['spec'] = V1SecurityOverviewSpec.constructFromObject(data['spec']);
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
     * @return {module:model/V1SecurityOverviewSpec}
     */
    getSpec() {
        return this.spec;
    }

    /**
     * @param {module:model/V1SecurityOverviewSpec} spec
     */
    setSpec(spec) {
        this['spec'] = spec;
    }

}

/**
 * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
 * @member {String} apiVersion
 */
V1SecurityOverview.prototype['apiVersion'] = undefined;

/**
 * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
 * @member {String} kind
 */
V1SecurityOverview.prototype['kind'] = undefined;

/**
 * @member {module:model/V1ObjectMeta} metadata
 */
V1SecurityOverview.prototype['metadata'] = undefined;

/**
 * @member {module:model/V1SecurityOverviewSpec} spec
 */
V1SecurityOverview.prototype['spec'] = undefined;






export default V1SecurityOverview;

