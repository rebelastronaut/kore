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
 * The V1Ownership model module.
 * @module model/V1Ownership
 * @version 0.0.1
 */
class V1Ownership {
    /**
     * Constructs a new <code>V1Ownership</code>.
     * @alias module:model/V1Ownership
     * @param group {String} 
     * @param kind {String} 
     * @param name {String} 
     * @param namespace {String} 
     * @param version {String} 
     */
    constructor(group, kind, name, namespace, version) { 
        
        V1Ownership.initialize(this, group, kind, name, namespace, version);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, group, kind, name, namespace, version) { 
        obj['group'] = group;
        obj['kind'] = kind;
        obj['name'] = name;
        obj['namespace'] = namespace;
        obj['version'] = version;
    }

    /**
     * Constructs a <code>V1Ownership</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Ownership} obj Optional instance to populate.
     * @return {module:model/V1Ownership} The populated <code>V1Ownership</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Ownership();

            if (data.hasOwnProperty('group')) {
                obj['group'] = ApiClient.convertToType(data['group'], 'String');
            }
            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('namespace')) {
                obj['namespace'] = ApiClient.convertToType(data['namespace'], 'String');
            }
            if (data.hasOwnProperty('version')) {
                obj['version'] = ApiClient.convertToType(data['version'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {String}
     */
    getGroup() {
        return this.group;
    }

    /**
     * @param {String} group
     */
    setGroup(group) {
        this['group'] = group;
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
    getName() {
        return this.name;
    }

    /**
     * @param {String} name
     */
    setName(name) {
        this['name'] = name;
    }
/**
     * @return {String}
     */
    getNamespace() {
        return this.namespace;
    }

    /**
     * @param {String} namespace
     */
    setNamespace(namespace) {
        this['namespace'] = namespace;
    }
/**
     * @return {String}
     */
    getVersion() {
        return this.version;
    }

    /**
     * @param {String} version
     */
    setVersion(version) {
        this['version'] = version;
    }

}

/**
 * @member {String} group
 */
V1Ownership.prototype['group'] = undefined;

/**
 * @member {String} kind
 */
V1Ownership.prototype['kind'] = undefined;

/**
 * @member {String} name
 */
V1Ownership.prototype['name'] = undefined;

/**
 * @member {String} namespace
 */
V1Ownership.prototype['namespace'] = undefined;

/**
 * @member {String} version
 */
V1Ownership.prototype['version'] = undefined;






export default V1Ownership;

