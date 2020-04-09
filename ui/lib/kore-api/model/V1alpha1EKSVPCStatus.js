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
import V1alpha1Infra from './V1alpha1Infra';

/**
 * The V1alpha1EKSVPCStatus model module.
 * @module model/V1alpha1EKSVPCStatus
 * @version 0.0.1
 */
class V1alpha1EKSVPCStatus {
    /**
     * Constructs a new <code>V1alpha1EKSVPCStatus</code>.
     * @alias module:model/V1alpha1EKSVPCStatus
     */
    constructor() { 
        
        V1alpha1EKSVPCStatus.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1alpha1EKSVPCStatus</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1alpha1EKSVPCStatus} obj Optional instance to populate.
     * @return {module:model/V1alpha1EKSVPCStatus} The populated <code>V1alpha1EKSVPCStatus</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1alpha1EKSVPCStatus();

            if (data.hasOwnProperty('conditions')) {
                obj['conditions'] = ApiClient.convertToType(data['conditions'], [V1Component]);
            }
            if (data.hasOwnProperty('infra')) {
                obj['infra'] = V1alpha1Infra.constructFromObject(data['infra']);
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {Array.<module:model/V1Component>}
     */
    getConditions() {
        return this.conditions;
    }

    /**
     * @param {Array.<module:model/V1Component>} conditions
     */
    setConditions(conditions) {
        this['conditions'] = conditions;
    }
/**
     * @return {module:model/V1alpha1Infra}
     */
    getInfra() {
        return this.infra;
    }

    /**
     * @param {module:model/V1alpha1Infra} infra
     */
    setInfra(infra) {
        this['infra'] = infra;
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
 * @member {Array.<module:model/V1Component>} conditions
 */
V1alpha1EKSVPCStatus.prototype['conditions'] = undefined;

/**
 * @member {module:model/V1alpha1Infra} infra
 */
V1alpha1EKSVPCStatus.prototype['infra'] = undefined;

/**
 * @member {String} status
 */
V1alpha1EKSVPCStatus.prototype['status'] = undefined;






export default V1alpha1EKSVPCStatus;

