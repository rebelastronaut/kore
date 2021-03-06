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
 * The V1TeamSpec model module.
 * @module model/V1TeamSpec
 * @version 0.0.1
 */
class V1TeamSpec {
    /**
     * Constructs a new <code>V1TeamSpec</code>.
     * @alias module:model/V1TeamSpec
     * @param description {String} 
     * @param summary {String} 
     */
    constructor(description, summary) { 
        
        V1TeamSpec.initialize(this, description, summary);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, description, summary) { 
        obj['description'] = description;
        obj['summary'] = summary;
    }

    /**
     * Constructs a <code>V1TeamSpec</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1TeamSpec} obj Optional instance to populate.
     * @return {module:model/V1TeamSpec} The populated <code>V1TeamSpec</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1TeamSpec();

            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
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
    getDescription() {
        return this.description;
    }

    /**
     * @param {String} description
     */
    setDescription(description) {
        this['description'] = description;
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
 * @member {String} description
 */
V1TeamSpec.prototype['description'] = undefined;

/**
 * @member {String} summary
 */
V1TeamSpec.prototype['summary'] = undefined;






export default V1TeamSpec;

