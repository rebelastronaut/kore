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
 * The V1beta1Region model module.
 * @module model/V1beta1Region
 * @version 0.0.1
 */
class V1beta1Region {
    /**
     * Constructs a new <code>V1beta1Region</code>.
     * @alias module:model/V1beta1Region
     * @param id {String} 
     * @param name {String} 
     */
    constructor(id, name) { 
        
        V1beta1Region.initialize(this, id, name);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, id, name) { 
        obj['id'] = id;
        obj['name'] = name;
    }

    /**
     * Constructs a <code>V1beta1Region</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1beta1Region} obj Optional instance to populate.
     * @return {module:model/V1beta1Region} The populated <code>V1beta1Region</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1beta1Region();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {String}
     */
    getId() {
        return this.id;
    }

    /**
     * @param {String} id
     */
    setId(id) {
        this['id'] = id;
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

}

/**
 * @member {String} id
 */
V1beta1Region.prototype['id'] = undefined;

/**
 * @member {String} name
 */
V1beta1Region.prototype['name'] = undefined;






export default V1beta1Region;

