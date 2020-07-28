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
import V1Component from './V1Component';

/**
 * The V1alpha1EKSNodeGroupStatus model module.
 * @module model/V1alpha1EKSNodeGroupStatus
 * @version 0.0.1
 */
class V1alpha1EKSNodeGroupStatus {
    /**
     * Constructs a new <code>V1alpha1EKSNodeGroupStatus</code>.
     * @alias module:model/V1alpha1EKSNodeGroupStatus
     */
    constructor() { 
        
        V1alpha1EKSNodeGroupStatus.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1alpha1EKSNodeGroupStatus</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1alpha1EKSNodeGroupStatus} obj Optional instance to populate.
     * @return {module:model/V1alpha1EKSNodeGroupStatus} The populated <code>V1alpha1EKSNodeGroupStatus</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1alpha1EKSNodeGroupStatus();

            if (data.hasOwnProperty('autoScalingGroupNames')) {
                obj['autoScalingGroupNames'] = ApiClient.convertToType(data['autoScalingGroupNames'], ['String']);
            }
            if (data.hasOwnProperty('conditions')) {
                obj['conditions'] = ApiClient.convertToType(data['conditions'], [V1Component]);
            }
            if (data.hasOwnProperty('nodeIAMRole')) {
                obj['nodeIAMRole'] = ApiClient.convertToType(data['nodeIAMRole'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {Array.<String>}
     */
    getAutoScalingGroupNames() {
        return this.autoScalingGroupNames;
    }

    /**
     * @param {Array.<String>} autoScalingGroupNames
     */
    setAutoScalingGroupNames(autoScalingGroupNames) {
        this['autoScalingGroupNames'] = autoScalingGroupNames;
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
     * @return {String}
     */
    getNodeIAMRole() {
        return this.nodeIAMRole;
    }

    /**
     * @param {String} nodeIAMRole
     */
    setNodeIAMRole(nodeIAMRole) {
        this['nodeIAMRole'] = nodeIAMRole;
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
 * @member {Array.<String>} autoScalingGroupNames
 */
V1alpha1EKSNodeGroupStatus.prototype['autoScalingGroupNames'] = undefined;

/**
 * @member {Array.<module:model/V1Component>} conditions
 */
V1alpha1EKSNodeGroupStatus.prototype['conditions'] = undefined;

/**
 * @member {String} nodeIAMRole
 */
V1alpha1EKSNodeGroupStatus.prototype['nodeIAMRole'] = undefined;

/**
 * @member {String} status
 */
V1alpha1EKSNodeGroupStatus.prototype['status'] = undefined;






export default V1alpha1EKSNodeGroupStatus;

