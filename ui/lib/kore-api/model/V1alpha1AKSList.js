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
import V1ListMeta from './V1ListMeta';
import V1alpha1AKS from './V1alpha1AKS';

/**
 * The V1alpha1AKSList model module.
 * @module model/V1alpha1AKSList
 * @version 0.0.1
 */
class V1alpha1AKSList {
    /**
     * Constructs a new <code>V1alpha1AKSList</code>.
     * @alias module:model/V1alpha1AKSList
     * @param items {Array.<module:model/V1alpha1AKS>} 
     */
    constructor(items) { 
        
        V1alpha1AKSList.initialize(this, items);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, items) { 
        obj['items'] = items;
    }

    /**
     * Constructs a <code>V1alpha1AKSList</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1alpha1AKSList} obj Optional instance to populate.
     * @return {module:model/V1alpha1AKSList} The populated <code>V1alpha1AKSList</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1alpha1AKSList();

            if (data.hasOwnProperty('apiVersion')) {
                obj['apiVersion'] = ApiClient.convertToType(data['apiVersion'], 'String');
            }
            if (data.hasOwnProperty('items')) {
                obj['items'] = ApiClient.convertToType(data['items'], [V1alpha1AKS]);
            }
            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('metadata')) {
                obj['metadata'] = V1ListMeta.constructFromObject(data['metadata']);
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
     * @return {Array.<module:model/V1alpha1AKS>}
     */
    getItems() {
        return this.items;
    }

    /**
     * @param {Array.<module:model/V1alpha1AKS>} items
     */
    setItems(items) {
        this['items'] = items;
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
     * @return {module:model/V1ListMeta}
     */
    getMetadata() {
        return this.metadata;
    }

    /**
     * @param {module:model/V1ListMeta} metadata
     */
    setMetadata(metadata) {
        this['metadata'] = metadata;
    }

}

/**
 * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
 * @member {String} apiVersion
 */
V1alpha1AKSList.prototype['apiVersion'] = undefined;

/**
 * @member {Array.<module:model/V1alpha1AKS>} items
 */
V1alpha1AKSList.prototype['items'] = undefined;

/**
 * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
 * @member {String} kind
 */
V1alpha1AKSList.prototype['kind'] = undefined;

/**
 * @member {module:model/V1ListMeta} metadata
 */
V1alpha1AKSList.prototype['metadata'] = undefined;






export default V1alpha1AKSList;

