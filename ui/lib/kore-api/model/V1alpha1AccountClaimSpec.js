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
 * The V1alpha1AccountClaimSpec model module.
 * @module model/V1alpha1AccountClaimSpec
 * @version 0.0.1
 */
class V1alpha1AccountClaimSpec {
    /**
     * Constructs a new <code>V1alpha1AccountClaimSpec</code>.
     * @alias module:model/V1alpha1AccountClaimSpec
     * @param accountName {String} 
     * @param organization {module:model/V1Ownership} 
     */
    constructor(accountName, organization) { 
        
        V1alpha1AccountClaimSpec.initialize(this, accountName, organization);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, accountName, organization) { 
        obj['accountName'] = accountName;
        obj['organization'] = organization;
    }

    /**
     * Constructs a <code>V1alpha1AccountClaimSpec</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1alpha1AccountClaimSpec} obj Optional instance to populate.
     * @return {module:model/V1alpha1AccountClaimSpec} The populated <code>V1alpha1AccountClaimSpec</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1alpha1AccountClaimSpec();

            if (data.hasOwnProperty('accountName')) {
                obj['accountName'] = ApiClient.convertToType(data['accountName'], 'String');
            }
            if (data.hasOwnProperty('organization')) {
                obj['organization'] = V1Ownership.constructFromObject(data['organization']);
            }
        }
        return obj;
    }

/**
     * @return {String}
     */
    getAccountName() {
        return this.accountName;
    }

    /**
     * @param {String} accountName
     */
    setAccountName(accountName) {
        this['accountName'] = accountName;
    }
/**
     * @return {module:model/V1Ownership}
     */
    getOrganization() {
        return this.organization;
    }

    /**
     * @param {module:model/V1Ownership} organization
     */
    setOrganization(organization) {
        this['organization'] = organization;
    }

}

/**
 * @member {String} accountName
 */
V1alpha1AccountClaimSpec.prototype['accountName'] = undefined;

/**
 * @member {module:model/V1Ownership} organization
 */
V1alpha1AccountClaimSpec.prototype['organization'] = undefined;






export default V1alpha1AccountClaimSpec;

