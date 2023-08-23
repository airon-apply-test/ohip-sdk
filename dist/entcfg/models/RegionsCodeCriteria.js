"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.RegionsCodeCriteriaToJSON = exports.RegionsCodeCriteriaFromJSONTyped = exports.RegionsCodeCriteriaFromJSON = exports.instanceOfRegionsCodeCriteria = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const RegionCodeType_1 = require("./RegionCodeType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the RegionsCodeCriteria interface.
 */
function instanceOfRegionsCodeCriteria(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRegionsCodeCriteria = instanceOfRegionsCodeCriteria;
function RegionsCodeCriteriaFromJSON(json) {
    return RegionsCodeCriteriaFromJSONTyped(json, false);
}
exports.RegionsCodeCriteriaFromJSON = RegionsCodeCriteriaFromJSON;
function RegionsCodeCriteriaFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'regionsCode': !(0, runtime_1.exists)(json, 'regionsCode') ? undefined : (json['regionsCode'].map(RegionCodeType_1.RegionCodeTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.RegionsCodeCriteriaFromJSONTyped = RegionsCodeCriteriaFromJSONTyped;
function RegionsCodeCriteriaToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'regionsCode': value.regionsCode === undefined ? undefined : (value.regionsCode.map(RegionCodeType_1.RegionCodeTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.RegionsCodeCriteriaToJSON = RegionsCodeCriteriaToJSON;
