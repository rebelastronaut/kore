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
 * The V1BasicAuth model module.
 * @module model/V1BasicAuth
 * @version 0.0.1
 */
class V1BasicAuth {
    /**
     * Constructs a new <code>V1BasicAuth</code>.
     * @alias module:model/V1BasicAuth
     * @param password {String} 
     */
    constructor(password) { 
        
        V1BasicAuth.initialize(this, password);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, password) { 
        obj['password'] = password;
    }

    /**
     * Constructs a <code>V1BasicAuth</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1BasicAuth} obj Optional instance to populate.
     * @return {module:model/V1BasicAuth} The populated <code>V1BasicAuth</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1BasicAuth();

            if (data.hasOwnProperty('password')) {
                obj['password'] = ApiClient.convertToType(data['password'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {String}
     */
    getPassword() {
        return this.password;
    }

    /**
     * @param {String} password
     */
    setPassword(password) {
        this['password'] = password;
    }

}

/**
 * @member {String} password
 */
V1BasicAuth.prototype['password'] = undefined;






export default V1BasicAuth;

