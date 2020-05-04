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
import V1Ownership from './V1Ownership';

/**
 * The V1ServiceSpec model module.
 * @module model/V1ServiceSpec
 * @version 0.0.1
 */
class V1ServiceSpec {
    /**
     * Constructs a new <code>V1ServiceSpec</code>.
     * @alias module:model/V1ServiceSpec
     * @param configuration {String} 
     * @param credentials {module:model/V1Ownership} 
     * @param kind {String} 
     * @param plan {String} 
     */
    constructor(configuration, credentials, kind, plan) { 
        
        V1ServiceSpec.initialize(this, configuration, credentials, kind, plan);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, configuration, credentials, kind, plan) { 
        obj['configuration'] = configuration;
        obj['credentials'] = credentials;
        obj['kind'] = kind;
        obj['plan'] = plan;
    }

    /**
     * Constructs a <code>V1ServiceSpec</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1ServiceSpec} obj Optional instance to populate.
     * @return {module:model/V1ServiceSpec} The populated <code>V1ServiceSpec</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1ServiceSpec();

            if (data.hasOwnProperty('configuration')) {
                obj['configuration'] = ApiClient.convertToType(data['configuration'], 'String');
            }
            if (data.hasOwnProperty('credentials')) {
                obj['credentials'] = V1Ownership.constructFromObject(data['credentials']);
            }
            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('plan')) {
                obj['plan'] = ApiClient.convertToType(data['plan'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {String}
     */
    getConfiguration() {
        return this.configuration;
    }

    /**
     * @param {String} configuration
     */
    setConfiguration(configuration) {
        this['configuration'] = configuration;
    }
/**
     * @return {module:model/V1Ownership}
     */
    getCredentials() {
        return this.credentials;
    }

    /**
     * @param {module:model/V1Ownership} credentials
     */
    setCredentials(credentials) {
        this['credentials'] = credentials;
    }
/**
     * @return {String}
     */
    getKind() {
        return this.kind;
    }

    /**
     * @param {String} kind
     */
    setKind(kind) {
        this['kind'] = kind;
    }
/**
     * @return {String}
     */
    getPlan() {
        return this.plan;
    }

    /**
     * @param {String} plan
     */
    setPlan(plan) {
        this['plan'] = plan;
    }

}

/**
 * @member {String} configuration
 */
V1ServiceSpec.prototype['configuration'] = undefined;

/**
 * @member {module:model/V1Ownership} credentials
 */
V1ServiceSpec.prototype['credentials'] = undefined;

/**
 * @member {String} kind
 */
V1ServiceSpec.prototype['kind'] = undefined;

/**
 * @member {String} plan
 */
V1ServiceSpec.prototype['plan'] = undefined;






export default V1ServiceSpec;

