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
import V1Ownership from './V1Ownership';

/**
 * The V1alpha1EKSVPCSpec model module.
 * @module model/V1alpha1EKSVPCSpec
 * @version 0.0.1
 */
class V1alpha1EKSVPCSpec {
    /**
     * Constructs a new <code>V1alpha1EKSVPCSpec</code>.
     * @alias module:model/V1alpha1EKSVPCSpec
     * @param credentials {module:model/V1Ownership} 
     * @param privateIPV4Cidr {String} 
     * @param region {String} 
     */
    constructor(credentials, privateIPV4Cidr, region) { 
        
        V1alpha1EKSVPCSpec.initialize(this, credentials, privateIPV4Cidr, region);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, credentials, privateIPV4Cidr, region) { 
        obj['credentials'] = credentials;
        obj['privateIPV4Cidr'] = privateIPV4Cidr;
        obj['region'] = region;
    }

    /**
     * Constructs a <code>V1alpha1EKSVPCSpec</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1alpha1EKSVPCSpec} obj Optional instance to populate.
     * @return {module:model/V1alpha1EKSVPCSpec} The populated <code>V1alpha1EKSVPCSpec</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1alpha1EKSVPCSpec();

            if (data.hasOwnProperty('cluster')) {
                obj['cluster'] = V1Ownership.constructFromObject(data['cluster']);
            }
            if (data.hasOwnProperty('credentials')) {
                obj['credentials'] = V1Ownership.constructFromObject(data['credentials']);
            }
            if (data.hasOwnProperty('privateIPV4Cidr')) {
                obj['privateIPV4Cidr'] = ApiClient.convertToType(data['privateIPV4Cidr'], 'String');
            }
            if (data.hasOwnProperty('region')) {
                obj['region'] = ApiClient.convertToType(data['region'], 'String');
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], {'String': 'String'});
            }
        }
        return obj;
    }

/**
     * @return {module:model/V1Ownership}
     */
    getCluster() {
        return this.cluster;
    }

    /**
     * @param {module:model/V1Ownership} cluster
     */
    setCluster(cluster) {
        this['cluster'] = cluster;
    }
/**
     * @return {module:model/V1Ownership}
     */
    getCredentials() {
        return this.credentials;
    }

    /**
     * @param {module:model/V1Ownership} credentials
     */
    setCredentials(credentials) {
        this['credentials'] = credentials;
    }
/**
     * @return {String}
     */
    getPrivateIPV4Cidr() {
        return this.privateIPV4Cidr;
    }

    /**
     * @param {String} privateIPV4Cidr
     */
    setPrivateIPV4Cidr(privateIPV4Cidr) {
        this['privateIPV4Cidr'] = privateIPV4Cidr;
    }
/**
     * @return {String}
     */
    getRegion() {
        return this.region;
    }

    /**
     * @param {String} region
     */
    setRegion(region) {
        this['region'] = region;
    }
/**
     * @return {Object.<String, String>}
     */
    getTags() {
        return this.tags;
    }

    /**
     * @param {Object.<String, String>} tags
     */
    setTags(tags) {
        this['tags'] = tags;
    }

}

/**
 * @member {module:model/V1Ownership} cluster
 */
V1alpha1EKSVPCSpec.prototype['cluster'] = undefined;

/**
 * @member {module:model/V1Ownership} credentials
 */
V1alpha1EKSVPCSpec.prototype['credentials'] = undefined;

/**
 * @member {String} privateIPV4Cidr
 */
V1alpha1EKSVPCSpec.prototype['privateIPV4Cidr'] = undefined;

/**
 * @member {String} region
 */
V1alpha1EKSVPCSpec.prototype['region'] = undefined;

/**
 * @member {Object.<String, String>} tags
 */
V1alpha1EKSVPCSpec.prototype['tags'] = undefined;






export default V1alpha1EKSVPCSpec;

