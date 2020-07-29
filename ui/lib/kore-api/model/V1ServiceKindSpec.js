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
 * The V1ServiceKindSpec model module.
 * @module model/V1ServiceKindSpec
 * @version 0.0.1
 */
class V1ServiceKindSpec {
    /**
     * Constructs a new <code>V1ServiceKindSpec</code>.
     * @alias module:model/V1ServiceKindSpec
     * @param enabled {Boolean} 
     * @param serviceAccessEnabled {Boolean} 
     * @param summary {String} 
     */
    constructor(enabled, serviceAccessEnabled, summary) { 
        
        V1ServiceKindSpec.initialize(this, enabled, serviceAccessEnabled, summary);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, enabled, serviceAccessEnabled, summary) { 
        obj['enabled'] = enabled;
        obj['serviceAccessEnabled'] = serviceAccessEnabled;
        obj['summary'] = summary;
    }

    /**
     * Constructs a <code>V1ServiceKindSpec</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1ServiceKindSpec} obj Optional instance to populate.
     * @return {module:model/V1ServiceKindSpec} The populated <code>V1ServiceKindSpec</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1ServiceKindSpec();

            if (data.hasOwnProperty('credentialSchema')) {
                obj['credentialSchema'] = ApiClient.convertToType(data['credentialSchema'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('displayName')) {
                obj['displayName'] = ApiClient.convertToType(data['displayName'], 'String');
            }
            if (data.hasOwnProperty('documentationURL')) {
                obj['documentationURL'] = ApiClient.convertToType(data['documentationURL'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('imageURL')) {
                obj['imageURL'] = ApiClient.convertToType(data['imageURL'], 'String');
            }
            if (data.hasOwnProperty('providerData')) {
                obj['providerData'] = ApiClient.convertToType(data['providerData'], 'String');
            }
            if (data.hasOwnProperty('schema')) {
                obj['schema'] = ApiClient.convertToType(data['schema'], 'String');
            }
            if (data.hasOwnProperty('serviceAccessEnabled')) {
                obj['serviceAccessEnabled'] = ApiClient.convertToType(data['serviceAccessEnabled'], 'Boolean');
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
    getCredentialSchema() {
        return this.credentialSchema;
    }

    /**
     * @param {String} credentialSchema
     */
    setCredentialSchema(credentialSchema) {
        this['credentialSchema'] = credentialSchema;
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
    getDisplayName() {
        return this.displayName;
    }

    /**
     * @param {String} displayName
     */
    setDisplayName(displayName) {
        this['displayName'] = displayName;
    }
/**
     * @return {String}
     */
    getDocumentationURL() {
        return this.documentationURL;
    }

    /**
     * @param {String} documentationURL
     */
    setDocumentationURL(documentationURL) {
        this['documentationURL'] = documentationURL;
    }
/**
     * @return {Boolean}
     */
    getEnabled() {
        return this.enabled;
    }

    /**
     * @param {Boolean} enabled
     */
    setEnabled(enabled) {
        this['enabled'] = enabled;
    }
/**
     * @return {String}
     */
    getImageURL() {
        return this.imageURL;
    }

    /**
     * @param {String} imageURL
     */
    setImageURL(imageURL) {
        this['imageURL'] = imageURL;
    }
/**
     * @return {String}
     */
    getProviderData() {
        return this.providerData;
    }

    /**
     * @param {String} providerData
     */
    setProviderData(providerData) {
        this['providerData'] = providerData;
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
/**
     * @return {Boolean}
     */
    getServiceAccessEnabled() {
        return this.serviceAccessEnabled;
    }

    /**
     * @param {Boolean} serviceAccessEnabled
     */
    setServiceAccessEnabled(serviceAccessEnabled) {
        this['serviceAccessEnabled'] = serviceAccessEnabled;
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
 * @member {String} credentialSchema
 */
V1ServiceKindSpec.prototype['credentialSchema'] = undefined;

/**
 * @member {String} description
 */
V1ServiceKindSpec.prototype['description'] = undefined;

/**
 * @member {String} displayName
 */
V1ServiceKindSpec.prototype['displayName'] = undefined;

/**
 * @member {String} documentationURL
 */
V1ServiceKindSpec.prototype['documentationURL'] = undefined;

/**
 * @member {Boolean} enabled
 */
V1ServiceKindSpec.prototype['enabled'] = undefined;

/**
 * @member {String} imageURL
 */
V1ServiceKindSpec.prototype['imageURL'] = undefined;

/**
 * @member {String} providerData
 */
V1ServiceKindSpec.prototype['providerData'] = undefined;

/**
 * @member {String} schema
 */
V1ServiceKindSpec.prototype['schema'] = undefined;

/**
 * @member {Boolean} serviceAccessEnabled
 */
V1ServiceKindSpec.prototype['serviceAccessEnabled'] = undefined;

/**
 * @member {String} summary
 */
V1ServiceKindSpec.prototype['summary'] = undefined;






export default V1ServiceKindSpec;

