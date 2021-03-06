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
 * The V1beta1AlertRuleSpec model module.
 * @module model/V1beta1AlertRuleSpec
 * @version 0.0.1
 */
class V1beta1AlertRuleSpec {
    /**
     * Constructs a new <code>V1beta1AlertRuleSpec</code>.
     * @alias module:model/V1beta1AlertRuleSpec
     * @param rawRule {String} 
     * @param resource {module:model/V1Ownership} 
     * @param severity {String} 
     * @param source {String} 
     * @param summary {String} 
     */
    constructor(rawRule, resource, severity, source, summary) { 
        
        V1beta1AlertRuleSpec.initialize(this, rawRule, resource, severity, source, summary);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, rawRule, resource, severity, source, summary) { 
        obj['rawRule'] = rawRule;
        obj['resource'] = resource;
        obj['severity'] = severity;
        obj['source'] = source;
        obj['summary'] = summary;
    }

    /**
     * Constructs a <code>V1beta1AlertRuleSpec</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1beta1AlertRuleSpec} obj Optional instance to populate.
     * @return {module:model/V1beta1AlertRuleSpec} The populated <code>V1beta1AlertRuleSpec</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1beta1AlertRuleSpec();

            if (data.hasOwnProperty('rawRule')) {
                obj['rawRule'] = ApiClient.convertToType(data['rawRule'], 'String');
            }
            if (data.hasOwnProperty('resource')) {
                obj['resource'] = V1Ownership.constructFromObject(data['resource']);
            }
            if (data.hasOwnProperty('ruleID')) {
                obj['ruleID'] = ApiClient.convertToType(data['ruleID'], 'String');
            }
            if (data.hasOwnProperty('severity')) {
                obj['severity'] = ApiClient.convertToType(data['severity'], 'String');
            }
            if (data.hasOwnProperty('source')) {
                obj['source'] = ApiClient.convertToType(data['source'], 'String');
            }
            if (data.hasOwnProperty('summary')) {
                obj['summary'] = ApiClient.convertToType(data['summary'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {String}
     */
    getRawRule() {
        return this.rawRule;
    }

    /**
     * @param {String} rawRule
     */
    setRawRule(rawRule) {
        this['rawRule'] = rawRule;
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
    getRuleID() {
        return this.ruleID;
    }

    /**
     * @param {String} ruleID
     */
    setRuleID(ruleID) {
        this['ruleID'] = ruleID;
    }
/**
     * @return {String}
     */
    getSeverity() {
        return this.severity;
    }

    /**
     * @param {String} severity
     */
    setSeverity(severity) {
        this['severity'] = severity;
    }
/**
     * @return {String}
     */
    getSource() {
        return this.source;
    }

    /**
     * @param {String} source
     */
    setSource(source) {
        this['source'] = source;
    }
/**
     * @return {String}
     */
    getSummary() {
        return this.summary;
    }

    /**
     * @param {String} summary
     */
    setSummary(summary) {
        this['summary'] = summary;
    }

}

/**
 * @member {String} rawRule
 */
V1beta1AlertRuleSpec.prototype['rawRule'] = undefined;

/**
 * @member {module:model/V1Ownership} resource
 */
V1beta1AlertRuleSpec.prototype['resource'] = undefined;

/**
 * @member {String} ruleID
 */
V1beta1AlertRuleSpec.prototype['ruleID'] = undefined;

/**
 * @member {String} severity
 */
V1beta1AlertRuleSpec.prototype['severity'] = undefined;

/**
 * @member {String} source
 */
V1beta1AlertRuleSpec.prototype['source'] = undefined;

/**
 * @member {String} summary
 */
V1beta1AlertRuleSpec.prototype['summary'] = undefined;






export default V1beta1AlertRuleSpec;

