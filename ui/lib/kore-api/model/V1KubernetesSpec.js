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
import V1ClusterUser from './V1ClusterUser';
import V1Ownership from './V1Ownership';

/**
 * The V1KubernetesSpec model module.
 * @module model/V1KubernetesSpec
 * @version 0.0.1
 */
class V1KubernetesSpec {
    /**
     * Constructs a new <code>V1KubernetesSpec</code>.
     * @alias module:model/V1KubernetesSpec
     */
    constructor() { 
        
        V1KubernetesSpec.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1KubernetesSpec</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1KubernetesSpec} obj Optional instance to populate.
     * @return {module:model/V1KubernetesSpec} The populated <code>V1KubernetesSpec</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1KubernetesSpec();

            if (data.hasOwnProperty('authProxyAllowedIPs')) {
                obj['authProxyAllowedIPs'] = ApiClient.convertToType(data['authProxyAllowedIPs'], ['String']);
            }
            if (data.hasOwnProperty('authProxyImage')) {
                obj['authProxyImage'] = ApiClient.convertToType(data['authProxyImage'], 'String');
            }
            if (data.hasOwnProperty('cluster')) {
                obj['cluster'] = V1Ownership.constructFromObject(data['cluster']);
            }
            if (data.hasOwnProperty('clusterUsers')) {
                obj['clusterUsers'] = ApiClient.convertToType(data['clusterUsers'], [V1ClusterUser]);
            }
            if (data.hasOwnProperty('defaultTeamRole')) {
                obj['defaultTeamRole'] = ApiClient.convertToType(data['defaultTeamRole'], 'String');
            }
            if (data.hasOwnProperty('domain')) {
                obj['domain'] = ApiClient.convertToType(data['domain'], 'String');
            }
            if (data.hasOwnProperty('enableDefaultTrafficBlock')) {
                obj['enableDefaultTrafficBlock'] = ApiClient.convertToType(data['enableDefaultTrafficBlock'], 'Boolean');
            }
            if (data.hasOwnProperty('inheritTeamMembers')) {
                obj['inheritTeamMembers'] = ApiClient.convertToType(data['inheritTeamMembers'], 'Boolean');
            }
            if (data.hasOwnProperty('provider')) {
                obj['provider'] = V1Ownership.constructFromObject(data['provider']);
            }
        }
        return obj;
    }

/**
     * @return {Array.<String>}
     */
    getAuthProxyAllowedIPs() {
        return this.authProxyAllowedIPs;
    }

    /**
     * @param {Array.<String>} authProxyAllowedIPs
     */
    setAuthProxyAllowedIPs(authProxyAllowedIPs) {
        this['authProxyAllowedIPs'] = authProxyAllowedIPs;
    }
/**
     * @return {String}
     */
    getAuthProxyImage() {
        return this.authProxyImage;
    }

    /**
     * @param {String} authProxyImage
     */
    setAuthProxyImage(authProxyImage) {
        this['authProxyImage'] = authProxyImage;
    }
/**
     * @return {module:model/V1Ownership}
     */
    getCluster() {
        return this.cluster;
    }

    /**
     * @param {module:model/V1Ownership} cluster
     */
    setCluster(cluster) {
        this['cluster'] = cluster;
    }
/**
     * @return {Array.<module:model/V1ClusterUser>}
     */
    getClusterUsers() {
        return this.clusterUsers;
    }

    /**
     * @param {Array.<module:model/V1ClusterUser>} clusterUsers
     */
    setClusterUsers(clusterUsers) {
        this['clusterUsers'] = clusterUsers;
    }
/**
     * @return {String}
     */
    getDefaultTeamRole() {
        return this.defaultTeamRole;
    }

    /**
     * @param {String} defaultTeamRole
     */
    setDefaultTeamRole(defaultTeamRole) {
        this['defaultTeamRole'] = defaultTeamRole;
    }
/**
     * @return {String}
     */
    getDomain() {
        return this.domain;
    }

    /**
     * @param {String} domain
     */
    setDomain(domain) {
        this['domain'] = domain;
    }
/**
     * @return {Boolean}
     */
    getEnableDefaultTrafficBlock() {
        return this.enableDefaultTrafficBlock;
    }

    /**
     * @param {Boolean} enableDefaultTrafficBlock
     */
    setEnableDefaultTrafficBlock(enableDefaultTrafficBlock) {
        this['enableDefaultTrafficBlock'] = enableDefaultTrafficBlock;
    }
/**
     * @return {Boolean}
     */
    getInheritTeamMembers() {
        return this.inheritTeamMembers;
    }

    /**
     * @param {Boolean} inheritTeamMembers
     */
    setInheritTeamMembers(inheritTeamMembers) {
        this['inheritTeamMembers'] = inheritTeamMembers;
    }
/**
     * @return {module:model/V1Ownership}
     */
    getProvider() {
        return this.provider;
    }

    /**
     * @param {module:model/V1Ownership} provider
     */
    setProvider(provider) {
        this['provider'] = provider;
    }

}

/**
 * @member {Array.<String>} authProxyAllowedIPs
 */
V1KubernetesSpec.prototype['authProxyAllowedIPs'] = undefined;

/**
 * @member {String} authProxyImage
 */
V1KubernetesSpec.prototype['authProxyImage'] = undefined;

/**
 * @member {module:model/V1Ownership} cluster
 */
V1KubernetesSpec.prototype['cluster'] = undefined;

/**
 * @member {Array.<module:model/V1ClusterUser>} clusterUsers
 */
V1KubernetesSpec.prototype['clusterUsers'] = undefined;

/**
 * @member {String} defaultTeamRole
 */
V1KubernetesSpec.prototype['defaultTeamRole'] = undefined;

/**
 * @member {String} domain
 */
V1KubernetesSpec.prototype['domain'] = undefined;

/**
 * @member {Boolean} enableDefaultTrafficBlock
 */
V1KubernetesSpec.prototype['enableDefaultTrafficBlock'] = undefined;

/**
 * @member {Boolean} inheritTeamMembers
 */
V1KubernetesSpec.prototype['inheritTeamMembers'] = undefined;

/**
 * @member {module:model/V1Ownership} provider
 */
V1KubernetesSpec.prototype['provider'] = undefined;






export default V1KubernetesSpec;

