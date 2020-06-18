/**
 * Appvia Kore API
 * Kore API provides the frontend API for the Appvia Kore (kore.appvia.io)
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
import V1Component from './V1Component';

/**
 * The V1ClusterStatus model module.
 * @module model/V1ClusterStatus
 * @version 0.0.1
 */
class V1ClusterStatus {
    /**
     * Constructs a new <code>V1ClusterStatus</code>.
     * @alias module:model/V1ClusterStatus
     */
    constructor() { 
        
        V1ClusterStatus.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1ClusterStatus</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1ClusterStatus} obj Optional instance to populate.
     * @return {module:model/V1ClusterStatus} The populated <code>V1ClusterStatus</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1ClusterStatus();

            if (data.hasOwnProperty('apiEndpoint')) {
                obj['apiEndpoint'] = ApiClient.convertToType(data['apiEndpoint'], 'String');
            }
            if (data.hasOwnProperty('authProxyEndpoint')) {
                obj['authProxyEndpoint'] = ApiClient.convertToType(data['authProxyEndpoint'], 'String');
            }
            if (data.hasOwnProperty('caCertificate')) {
                obj['caCertificate'] = ApiClient.convertToType(data['caCertificate'], 'String');
            }
            if (data.hasOwnProperty('components')) {
                obj['components'] = ApiClient.convertToType(data['components'], [V1Component]);
            }
            if (data.hasOwnProperty('message')) {
                obj['message'] = ApiClient.convertToType(data['message'], 'String');
            }
            if (data.hasOwnProperty('providerData')) {
                obj['providerData'] = ApiClient.convertToType(data['providerData'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {String}
     */
    getApiEndpoint() {
        return this.apiEndpoint;
    }

    /**
     * @param {String} apiEndpoint
     */
    setApiEndpoint(apiEndpoint) {
        this['apiEndpoint'] = apiEndpoint;
    }
/**
     * @return {String}
     */
    getAuthProxyEndpoint() {
        return this.authProxyEndpoint;
    }

    /**
     * @param {String} authProxyEndpoint
     */
    setAuthProxyEndpoint(authProxyEndpoint) {
        this['authProxyEndpoint'] = authProxyEndpoint;
    }
/**
     * @return {String}
     */
    getCaCertificate() {
        return this.caCertificate;
    }

    /**
     * @param {String} caCertificate
     */
    setCaCertificate(caCertificate) {
        this['caCertificate'] = caCertificate;
    }
/**
     * @return {Array.<module:model/V1Component>}
     */
    getComponents() {
        return this.components;
    }

    /**
     * @param {Array.<module:model/V1Component>} components
     */
    setComponents(components) {
        this['components'] = components;
    }
/**
     * @return {String}
     */
    getMessage() {
        return this.message;
    }

    /**
     * @param {String} message
     */
    setMessage(message) {
        this['message'] = message;
    }
/**
     * @return {String}
     */
    getProviderData() {
        return this.providerData;
    }

    /**
     * @param {String} providerData
     */
    setProviderData(providerData) {
        this['providerData'] = providerData;
    }
/**
     * @return {String}
     */
    getStatus() {
        return this.status;
    }

    /**
     * @param {String} status
     */
    setStatus(status) {
        this['status'] = status;
    }

}

/**
 * @member {String} apiEndpoint
 */
V1ClusterStatus.prototype['apiEndpoint'] = undefined;

/**
 * @member {String} authProxyEndpoint
 */
V1ClusterStatus.prototype['authProxyEndpoint'] = undefined;

/**
 * @member {String} caCertificate
 */
V1ClusterStatus.prototype['caCertificate'] = undefined;

/**
 * @member {Array.<module:model/V1Component>} components
 */
V1ClusterStatus.prototype['components'] = undefined;

/**
 * @member {String} message
 */
V1ClusterStatus.prototype['message'] = undefined;

/**
 * @member {String} providerData
 */
V1ClusterStatus.prototype['providerData'] = undefined;

/**
 * @member {String} status
 */
V1ClusterStatus.prototype['status'] = undefined;






export default V1ClusterStatus;

