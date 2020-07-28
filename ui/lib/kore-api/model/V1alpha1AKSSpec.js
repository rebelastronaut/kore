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
import V1alpha1AgentPoolProfile from './V1alpha1AgentPoolProfile';
import V1alpha1LinuxProfile from './V1alpha1LinuxProfile';
import V1alpha1WindowsProfile from './V1alpha1WindowsProfile';

/**
 * The V1alpha1AKSSpec model module.
 * @module model/V1alpha1AKSSpec
 * @version 0.0.1
 */
class V1alpha1AKSSpec {
    /**
     * Constructs a new <code>V1alpha1AKSSpec</code>.
     * @alias module:model/V1alpha1AKSSpec
     * @param agentPoolProfiles {Array.<module:model/V1alpha1AgentPoolProfile>} 
     * @param credentials {module:model/V1Ownership} 
     * @param description {String} 
     * @param dnsPrefix {String} 
     * @param location {String} 
     * @param networkPlugin {String} 
     */
    constructor(agentPoolProfiles, credentials, description, dnsPrefix, location, networkPlugin) { 
        
        V1alpha1AKSSpec.initialize(this, agentPoolProfiles, credentials, description, dnsPrefix, location, networkPlugin);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, agentPoolProfiles, credentials, description, dnsPrefix, location, networkPlugin) { 
        obj['agentPoolProfiles'] = agentPoolProfiles;
        obj['credentials'] = credentials;
        obj['description'] = description;
        obj['dnsPrefix'] = dnsPrefix;
        obj['location'] = location;
        obj['networkPlugin'] = networkPlugin;
    }

    /**
     * Constructs a <code>V1alpha1AKSSpec</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1alpha1AKSSpec} obj Optional instance to populate.
     * @return {module:model/V1alpha1AKSSpec} The populated <code>V1alpha1AKSSpec</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1alpha1AKSSpec();

            if (data.hasOwnProperty('agentPoolProfiles')) {
                obj['agentPoolProfiles'] = ApiClient.convertToType(data['agentPoolProfiles'], [V1alpha1AgentPoolProfile]);
            }
            if (data.hasOwnProperty('authorizedIPRanges')) {
                obj['authorizedIPRanges'] = ApiClient.convertToType(data['authorizedIPRanges'], ['String']);
            }
            if (data.hasOwnProperty('cluster')) {
                obj['cluster'] = V1Ownership.constructFromObject(data['cluster']);
            }
            if (data.hasOwnProperty('credentials')) {
                obj['credentials'] = V1Ownership.constructFromObject(data['credentials']);
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('dnsPrefix')) {
                obj['dnsPrefix'] = ApiClient.convertToType(data['dnsPrefix'], 'String');
            }
            if (data.hasOwnProperty('enablePodSecurityPolicy')) {
                obj['enablePodSecurityPolicy'] = ApiClient.convertToType(data['enablePodSecurityPolicy'], 'Boolean');
            }
            if (data.hasOwnProperty('enablePrivateCluster')) {
                obj['enablePrivateCluster'] = ApiClient.convertToType(data['enablePrivateCluster'], 'Boolean');
            }
            if (data.hasOwnProperty('kubernetesVersion')) {
                obj['kubernetesVersion'] = ApiClient.convertToType(data['kubernetesVersion'], 'String');
            }
            if (data.hasOwnProperty('linuxProfile')) {
                obj['linuxProfile'] = V1alpha1LinuxProfile.constructFromObject(data['linuxProfile']);
            }
            if (data.hasOwnProperty('location')) {
                obj['location'] = ApiClient.convertToType(data['location'], 'String');
            }
            if (data.hasOwnProperty('networkPlugin')) {
                obj['networkPlugin'] = ApiClient.convertToType(data['networkPlugin'], 'String');
            }
            if (data.hasOwnProperty('networkPolicy')) {
                obj['networkPolicy'] = ApiClient.convertToType(data['networkPolicy'], 'String');
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], {'String': 'String'});
            }
            if (data.hasOwnProperty('windowsProfile')) {
                obj['windowsProfile'] = V1alpha1WindowsProfile.constructFromObject(data['windowsProfile']);
            }
        }
        return obj;
    }

/**
     * @return {Array.<module:model/V1alpha1AgentPoolProfile>}
     */
    getAgentPoolProfiles() {
        return this.agentPoolProfiles;
    }

    /**
     * @param {Array.<module:model/V1alpha1AgentPoolProfile>} agentPoolProfiles
     */
    setAgentPoolProfiles(agentPoolProfiles) {
        this['agentPoolProfiles'] = agentPoolProfiles;
    }
/**
     * @return {Array.<String>}
     */
    getAuthorizedIPRanges() {
        return this.authorizedIPRanges;
    }

    /**
     * @param {Array.<String>} authorizedIPRanges
     */
    setAuthorizedIPRanges(authorizedIPRanges) {
        this['authorizedIPRanges'] = authorizedIPRanges;
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
    getDnsPrefix() {
        return this.dnsPrefix;
    }

    /**
     * @param {String} dnsPrefix
     */
    setDnsPrefix(dnsPrefix) {
        this['dnsPrefix'] = dnsPrefix;
    }
/**
     * @return {Boolean}
     */
    getEnablePodSecurityPolicy() {
        return this.enablePodSecurityPolicy;
    }

    /**
     * @param {Boolean} enablePodSecurityPolicy
     */
    setEnablePodSecurityPolicy(enablePodSecurityPolicy) {
        this['enablePodSecurityPolicy'] = enablePodSecurityPolicy;
    }
/**
     * @return {Boolean}
     */
    getEnablePrivateCluster() {
        return this.enablePrivateCluster;
    }

    /**
     * @param {Boolean} enablePrivateCluster
     */
    setEnablePrivateCluster(enablePrivateCluster) {
        this['enablePrivateCluster'] = enablePrivateCluster;
    }
/**
     * @return {String}
     */
    getKubernetesVersion() {
        return this.kubernetesVersion;
    }

    /**
     * @param {String} kubernetesVersion
     */
    setKubernetesVersion(kubernetesVersion) {
        this['kubernetesVersion'] = kubernetesVersion;
    }
/**
     * @return {module:model/V1alpha1LinuxProfile}
     */
    getLinuxProfile() {
        return this.linuxProfile;
    }

    /**
     * @param {module:model/V1alpha1LinuxProfile} linuxProfile
     */
    setLinuxProfile(linuxProfile) {
        this['linuxProfile'] = linuxProfile;
    }
/**
     * @return {String}
     */
    getLocation() {
        return this.location;
    }

    /**
     * @param {String} location
     */
    setLocation(location) {
        this['location'] = location;
    }
/**
     * @return {String}
     */
    getNetworkPlugin() {
        return this.networkPlugin;
    }

    /**
     * @param {String} networkPlugin
     */
    setNetworkPlugin(networkPlugin) {
        this['networkPlugin'] = networkPlugin;
    }
/**
     * @return {String}
     */
    getNetworkPolicy() {
        return this.networkPolicy;
    }

    /**
     * @param {String} networkPolicy
     */
    setNetworkPolicy(networkPolicy) {
        this['networkPolicy'] = networkPolicy;
    }
/**
     * @return {Object.<String, String>}
     */
    getTags() {
        return this.tags;
    }

    /**
     * @param {Object.<String, String>} tags
     */
    setTags(tags) {
        this['tags'] = tags;
    }
/**
     * @return {module:model/V1alpha1WindowsProfile}
     */
    getWindowsProfile() {
        return this.windowsProfile;
    }

    /**
     * @param {module:model/V1alpha1WindowsProfile} windowsProfile
     */
    setWindowsProfile(windowsProfile) {
        this['windowsProfile'] = windowsProfile;
    }

}

/**
 * @member {Array.<module:model/V1alpha1AgentPoolProfile>} agentPoolProfiles
 */
V1alpha1AKSSpec.prototype['agentPoolProfiles'] = undefined;

/**
 * @member {Array.<String>} authorizedIPRanges
 */
V1alpha1AKSSpec.prototype['authorizedIPRanges'] = undefined;

/**
 * @member {module:model/V1Ownership} cluster
 */
V1alpha1AKSSpec.prototype['cluster'] = undefined;

/**
 * @member {module:model/V1Ownership} credentials
 */
V1alpha1AKSSpec.prototype['credentials'] = undefined;

/**
 * @member {String} description
 */
V1alpha1AKSSpec.prototype['description'] = undefined;

/**
 * @member {String} dnsPrefix
 */
V1alpha1AKSSpec.prototype['dnsPrefix'] = undefined;

/**
 * @member {Boolean} enablePodSecurityPolicy
 */
V1alpha1AKSSpec.prototype['enablePodSecurityPolicy'] = undefined;

/**
 * @member {Boolean} enablePrivateCluster
 */
V1alpha1AKSSpec.prototype['enablePrivateCluster'] = undefined;

/**
 * @member {String} kubernetesVersion
 */
V1alpha1AKSSpec.prototype['kubernetesVersion'] = undefined;

/**
 * @member {module:model/V1alpha1LinuxProfile} linuxProfile
 */
V1alpha1AKSSpec.prototype['linuxProfile'] = undefined;

/**
 * @member {String} location
 */
V1alpha1AKSSpec.prototype['location'] = undefined;

/**
 * @member {String} networkPlugin
 */
V1alpha1AKSSpec.prototype['networkPlugin'] = undefined;

/**
 * @member {String} networkPolicy
 */
V1alpha1AKSSpec.prototype['networkPolicy'] = undefined;

/**
 * @member {Object.<String, String>} tags
 */
V1alpha1AKSSpec.prototype['tags'] = undefined;

/**
 * @member {module:model/V1alpha1WindowsProfile} windowsProfile
 */
V1alpha1AKSSpec.prototype['windowsProfile'] = undefined;






export default V1alpha1AKSSpec;

