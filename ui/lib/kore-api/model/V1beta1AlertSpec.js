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

/**
 * The V1beta1AlertSpec model module.
 * @module model/V1beta1AlertSpec
 * @version 0.0.1
 */
class V1beta1AlertSpec {
    /**
     * Constructs a new <code>V1beta1AlertSpec</code>.
     * @alias module:model/V1beta1AlertSpec
     * @param summary {String} 
     */
    constructor(summary) { 
        
        V1beta1AlertSpec.initialize(this, summary);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, summary) { 
        obj['summary'] = summary;
    }

    /**
     * Constructs a <code>V1beta1AlertSpec</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1beta1AlertSpec} obj Optional instance to populate.
     * @return {module:model/V1beta1AlertSpec} The populated <code>V1beta1AlertSpec</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1beta1AlertSpec();

            if (data.hasOwnProperty('alertID')) {
                obj['alertID'] = ApiClient.convertToType(data['alertID'], 'String');
            }
            if (data.hasOwnProperty('event')) {
                obj['event'] = ApiClient.convertToType(data['event'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], {'String': 'String'});
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
    getAlertID() {
        return this.alertID;
    }

    /**
     * @param {String} alertID
     */
    setAlertID(alertID) {
        this['alertID'] = alertID;
    }
/**
     * @return {String}
     */
    getEvent() {
        return this.event;
    }

    /**
     * @param {String} event
     */
    setEvent(event) {
        this['event'] = event;
    }
/**
     * @return {Object.<String, String>}
     */
    getLabels() {
        return this.labels;
    }

    /**
     * @param {Object.<String, String>} labels
     */
    setLabels(labels) {
        this['labels'] = labels;
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
 * @member {String} alertID
 */
V1beta1AlertSpec.prototype['alertID'] = undefined;

/**
 * @member {String} event
 */
V1beta1AlertSpec.prototype['event'] = undefined;

/**
 * @member {Object.<String, String>} labels
 */
V1beta1AlertSpec.prototype['labels'] = undefined;

/**
 * @member {String} summary
 */
V1beta1AlertSpec.prototype['summary'] = undefined;






export default V1beta1AlertSpec;

