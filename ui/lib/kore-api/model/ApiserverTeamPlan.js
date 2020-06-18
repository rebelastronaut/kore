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
import V1PlanSpec from './V1PlanSpec';

/**
 * The ApiserverTeamPlan model module.
 * @module model/ApiserverTeamPlan
 * @version 0.0.1
 */
class ApiserverTeamPlan {
    /**
     * Constructs a new <code>ApiserverTeamPlan</code>.
     * @alias module:model/ApiserverTeamPlan
     * @param editableParams {Array.<String>} 
     * @param schema {String} 
     */
    constructor(editableParams, schema) { 
        
        ApiserverTeamPlan.initialize(this, editableParams, schema);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, editableParams, schema) { 
        obj['editableParams'] = editableParams;
        obj['schema'] = schema;
    }

    /**
     * Constructs a <code>ApiserverTeamPlan</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiserverTeamPlan} obj Optional instance to populate.
     * @return {module:model/ApiserverTeamPlan} The populated <code>ApiserverTeamPlan</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiserverTeamPlan();

            if (data.hasOwnProperty('editableParams')) {
                obj['editableParams'] = ApiClient.convertToType(data['editableParams'], ['String']);
            }
            if (data.hasOwnProperty('plan')) {
                obj['plan'] = V1PlanSpec.constructFromObject(data['plan']);
            }
            if (data.hasOwnProperty('schema')) {
                obj['schema'] = ApiClient.convertToType(data['schema'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {Array.<String>}
     */
    getEditableParams() {
        return this.editableParams;
    }

    /**
     * @param {Array.<String>} editableParams
     */
    setEditableParams(editableParams) {
        this['editableParams'] = editableParams;
    }
/**
     * @return {module:model/V1PlanSpec}
     */
    getPlan() {
        return this.plan;
    }

    /**
     * @param {module:model/V1PlanSpec} plan
     */
    setPlan(plan) {
        this['plan'] = plan;
    }
/**
     * @return {String}
     */
    getSchema() {
        return this.schema;
    }

    /**
     * @param {String} schema
     */
    setSchema(schema) {
        this['schema'] = schema;
    }

}

/**
 * @member {Array.<String>} editableParams
 */
ApiserverTeamPlan.prototype['editableParams'] = undefined;

/**
 * @member {module:model/V1PlanSpec} plan
 */
ApiserverTeamPlan.prototype['plan'] = undefined;

/**
 * @member {String} schema
 */
ApiserverTeamPlan.prototype['schema'] = undefined;






export default ApiserverTeamPlan;

