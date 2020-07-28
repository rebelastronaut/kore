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
import V1alpha1NodeTaint from './V1alpha1NodeTaint';

/**
 * The V1alpha1GKENodePool model module.
 * @module model/V1alpha1GKENodePool
 * @version 0.0.1
 */
class V1alpha1GKENodePool {
    /**
     * Constructs a new <code>V1alpha1GKENodePool</code>.
     * @alias module:model/V1alpha1GKENodePool
     * @param diskSize {Number} 
     * @param enableAutorepair {Boolean} 
     * @param enableAutoscaler {Boolean} 
     * @param enableAutoupgrade {Boolean} 
     * @param imageType {String} 
     * @param machineType {String} 
     * @param maxPodsPerNode {Number} 
     * @param maxSize {Number} 
     * @param minSize {Number} 
     * @param name {String} 
     * @param preemptible {Boolean} 
     * @param size {Number} 
     * @param version {String} 
     */
    constructor(diskSize, enableAutorepair, enableAutoscaler, enableAutoupgrade, imageType, machineType, maxPodsPerNode, maxSize, minSize, name, preemptible, size, version) { 
        
        V1alpha1GKENodePool.initialize(this, diskSize, enableAutorepair, enableAutoscaler, enableAutoupgrade, imageType, machineType, maxPodsPerNode, maxSize, minSize, name, preemptible, size, version);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, diskSize, enableAutorepair, enableAutoscaler, enableAutoupgrade, imageType, machineType, maxPodsPerNode, maxSize, minSize, name, preemptible, size, version) { 
        obj['diskSize'] = diskSize;
        obj['enableAutorepair'] = enableAutorepair;
        obj['enableAutoscaler'] = enableAutoscaler;
        obj['enableAutoupgrade'] = enableAutoupgrade;
        obj['imageType'] = imageType;
        obj['machineType'] = machineType;
        obj['maxPodsPerNode'] = maxPodsPerNode;
        obj['maxSize'] = maxSize;
        obj['minSize'] = minSize;
        obj['name'] = name;
        obj['preemptible'] = preemptible;
        obj['size'] = size;
        obj['version'] = version;
    }

    /**
     * Constructs a <code>V1alpha1GKENodePool</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1alpha1GKENodePool} obj Optional instance to populate.
     * @return {module:model/V1alpha1GKENodePool} The populated <code>V1alpha1GKENodePool</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1alpha1GKENodePool();

            if (data.hasOwnProperty('diskSize')) {
                obj['diskSize'] = ApiClient.convertToType(data['diskSize'], 'Number');
            }
            if (data.hasOwnProperty('enableAutorepair')) {
                obj['enableAutorepair'] = ApiClient.convertToType(data['enableAutorepair'], 'Boolean');
            }
            if (data.hasOwnProperty('enableAutoscaler')) {
                obj['enableAutoscaler'] = ApiClient.convertToType(data['enableAutoscaler'], 'Boolean');
            }
            if (data.hasOwnProperty('enableAutoupgrade')) {
                obj['enableAutoupgrade'] = ApiClient.convertToType(data['enableAutoupgrade'], 'Boolean');
            }
            if (data.hasOwnProperty('imageType')) {
                obj['imageType'] = ApiClient.convertToType(data['imageType'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], {'String': 'String'});
            }
            if (data.hasOwnProperty('machineType')) {
                obj['machineType'] = ApiClient.convertToType(data['machineType'], 'String');
            }
            if (data.hasOwnProperty('maxPodsPerNode')) {
                obj['maxPodsPerNode'] = ApiClient.convertToType(data['maxPodsPerNode'], 'Number');
            }
            if (data.hasOwnProperty('maxSize')) {
                obj['maxSize'] = ApiClient.convertToType(data['maxSize'], 'Number');
            }
            if (data.hasOwnProperty('minSize')) {
                obj['minSize'] = ApiClient.convertToType(data['minSize'], 'Number');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('preemptible')) {
                obj['preemptible'] = ApiClient.convertToType(data['preemptible'], 'Boolean');
            }
            if (data.hasOwnProperty('size')) {
                obj['size'] = ApiClient.convertToType(data['size'], 'Number');
            }
            if (data.hasOwnProperty('taints')) {
                obj['taints'] = ApiClient.convertToType(data['taints'], [V1alpha1NodeTaint]);
            }
            if (data.hasOwnProperty('version')) {
                obj['version'] = ApiClient.convertToType(data['version'], 'String');
            }
        }
        return obj;
    }

/**
     * @return {Number}
     */
    getDiskSize() {
        return this.diskSize;
    }

    /**
     * @param {Number} diskSize
     */
    setDiskSize(diskSize) {
        this['diskSize'] = diskSize;
    }
/**
     * @return {Boolean}
     */
    getEnableAutorepair() {
        return this.enableAutorepair;
    }

    /**
     * @param {Boolean} enableAutorepair
     */
    setEnableAutorepair(enableAutorepair) {
        this['enableAutorepair'] = enableAutorepair;
    }
/**
     * @return {Boolean}
     */
    getEnableAutoscaler() {
        return this.enableAutoscaler;
    }

    /**
     * @param {Boolean} enableAutoscaler
     */
    setEnableAutoscaler(enableAutoscaler) {
        this['enableAutoscaler'] = enableAutoscaler;
    }
/**
     * @return {Boolean}
     */
    getEnableAutoupgrade() {
        return this.enableAutoupgrade;
    }

    /**
     * @param {Boolean} enableAutoupgrade
     */
    setEnableAutoupgrade(enableAutoupgrade) {
        this['enableAutoupgrade'] = enableAutoupgrade;
    }
/**
     * @return {String}
     */
    getImageType() {
        return this.imageType;
    }

    /**
     * @param {String} imageType
     */
    setImageType(imageType) {
        this['imageType'] = imageType;
    }
/**
     * @return {Object.<String, String>}
     */
    getLabels() {
        return this.labels;
    }

    /**
     * @param {Object.<String, String>} labels
     */
    setLabels(labels) {
        this['labels'] = labels;
    }
/**
     * @return {String}
     */
    getMachineType() {
        return this.machineType;
    }

    /**
     * @param {String} machineType
     */
    setMachineType(machineType) {
        this['machineType'] = machineType;
    }
/**
     * @return {Number}
     */
    getMaxPodsPerNode() {
        return this.maxPodsPerNode;
    }

    /**
     * @param {Number} maxPodsPerNode
     */
    setMaxPodsPerNode(maxPodsPerNode) {
        this['maxPodsPerNode'] = maxPodsPerNode;
    }
/**
     * @return {Number}
     */
    getMaxSize() {
        return this.maxSize;
    }

    /**
     * @param {Number} maxSize
     */
    setMaxSize(maxSize) {
        this['maxSize'] = maxSize;
    }
/**
     * @return {Number}
     */
    getMinSize() {
        return this.minSize;
    }

    /**
     * @param {Number} minSize
     */
    setMinSize(minSize) {
        this['minSize'] = minSize;
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
     * @return {Boolean}
     */
    getPreemptible() {
        return this.preemptible;
    }

    /**
     * @param {Boolean} preemptible
     */
    setPreemptible(preemptible) {
        this['preemptible'] = preemptible;
    }
/**
     * @return {Number}
     */
    getSize() {
        return this.size;
    }

    /**
     * @param {Number} size
     */
    setSize(size) {
        this['size'] = size;
    }
/**
     * @return {Array.<module:model/V1alpha1NodeTaint>}
     */
    getTaints() {
        return this.taints;
    }

    /**
     * @param {Array.<module:model/V1alpha1NodeTaint>} taints
     */
    setTaints(taints) {
        this['taints'] = taints;
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
 * @member {Number} diskSize
 */
V1alpha1GKENodePool.prototype['diskSize'] = undefined;

/**
 * @member {Boolean} enableAutorepair
 */
V1alpha1GKENodePool.prototype['enableAutorepair'] = undefined;

/**
 * @member {Boolean} enableAutoscaler
 */
V1alpha1GKENodePool.prototype['enableAutoscaler'] = undefined;

/**
 * @member {Boolean} enableAutoupgrade
 */
V1alpha1GKENodePool.prototype['enableAutoupgrade'] = undefined;

/**
 * @member {String} imageType
 */
V1alpha1GKENodePool.prototype['imageType'] = undefined;

/**
 * @member {Object.<String, String>} labels
 */
V1alpha1GKENodePool.prototype['labels'] = undefined;

/**
 * @member {String} machineType
 */
V1alpha1GKENodePool.prototype['machineType'] = undefined;

/**
 * @member {Number} maxPodsPerNode
 */
V1alpha1GKENodePool.prototype['maxPodsPerNode'] = undefined;

/**
 * @member {Number} maxSize
 */
V1alpha1GKENodePool.prototype['maxSize'] = undefined;

/**
 * @member {Number} minSize
 */
V1alpha1GKENodePool.prototype['minSize'] = undefined;

/**
 * @member {String} name
 */
V1alpha1GKENodePool.prototype['name'] = undefined;

/**
 * @member {Boolean} preemptible
 */
V1alpha1GKENodePool.prototype['preemptible'] = undefined;

/**
 * @member {Number} size
 */
V1alpha1GKENodePool.prototype['size'] = undefined;

/**
 * @member {Array.<module:model/V1alpha1NodeTaint>} taints
 */
V1alpha1GKENodePool.prototype['taints'] = undefined;

/**
 * @member {String} version
 */
V1alpha1GKENodePool.prototype['version'] = undefined;






export default V1alpha1GKENodePool;

