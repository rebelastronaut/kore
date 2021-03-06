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
 * The TypesIssuedToken model module.
 * @module model/TypesIssuedToken
 * @version 0.0.1
 */
class TypesIssuedToken {
    /**
     * Constructs a new <code>TypesIssuedToken</code>.
     * @alias module:model/TypesIssuedToken
     */
    constructor() { 
        
        TypesIssuedToken.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TypesIssuedToken</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TypesIssuedToken} obj Optional instance to populate.
     * @return {module:model/TypesIssuedToken} The populated <code>TypesIssuedToken</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TypesIssuedToken();

            if (data.hasOwnProperty('expires')) {
                obj['expires'] = ApiClient.convertToType(data['expires'], 'Number');
            }
            if (data.hasOwnProperty('refreshToken')) {
                obj['refreshToken'] = ApiClient.convertToType(data['refreshToken'], 'String');
            }
            if (data.hasOwnProperty('token')) {
                obj['token'] = ApiClient.convertToType(data['token'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {Number}
     */
    getExpires() {
        return this.expires;
    }

    /**
     * @param {Number} expires
     */
    setExpires(expires) {
        this['expires'] = expires;
    }
/**
     * @return {String}
     */
    getRefreshToken() {
        return this.refreshToken;
    }

    /**
     * @param {String} refreshToken
     */
    setRefreshToken(refreshToken) {
        this['refreshToken'] = refreshToken;
    }
/**
     * @return {String}
     */
    getToken() {
        return this.token;
    }

    /**
     * @param {String} token
     */
    setToken(token) {
        this['token'] = token;
    }

}

/**
 * @member {Number} expires
 */
TypesIssuedToken.prototype['expires'] = undefined;

/**
 * @member {String} refreshToken
 */
TypesIssuedToken.prototype['refreshToken'] = undefined;

/**
 * @member {String} token
 */
TypesIssuedToken.prototype['token'] = undefined;






export default TypesIssuedToken;

