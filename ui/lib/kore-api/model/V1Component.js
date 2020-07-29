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
 * The V1Component model module.
 * @module model/V1Component
 * @version 0.0.1
 */
class V1Component {
    /**
     * Constructs a new <code>V1Component</code>.
     * @alias module:model/V1Component
     */
    constructor() { 
        
        V1Component.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1Component</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Component} obj Optional instance to populate.
     * @return {module:model/V1Component} The populated <code>V1Component</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Component();

            if (data.hasOwnProperty('detail')) {
                obj['detail'] = ApiClient.convertToType(data['detail'], 'String');
            }
            if (data.hasOwnProperty('message')) {
                obj['message'] = ApiClient.convertToType(data['message'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('resource')) {
                obj['resource'] = V1Ownership.constructFromObject(data['resource']);
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
    getDetail() {
        return this.detail;
    }

    /**
     * @param {String} detail
     */
    setDetail(detail) {
        this['detail'] = detail;
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
    getName() {
        return this.name;
    }

    /**
     * @param {String} name
     */
    setName(name) {
        this['name'] = name;
    }
/**
     * @return {module:model/V1Ownership}
     */
    getResource() {
        return this.resource;
    }

    /**
     * @param {module:model/V1Ownership} resource
     */
    setResource(resource) {
        this['resource'] = resource;
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
 * @member {String} detail
 */
V1Component.prototype['detail'] = undefined;

/**
 * @member {String} message
 */
V1Component.prototype['message'] = undefined;

/**
 * @member {String} name
 */
V1Component.prototype['name'] = undefined;

/**
 * @member {module:model/V1Ownership} resource
 */
V1Component.prototype['resource'] = undefined;

/**
 * @member {String} status
 */
V1Component.prototype['status'] = undefined;






export default V1Component;

