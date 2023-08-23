"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Integration Configuration API
 * APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.TemplateDeviceLocationsCriteriaToJSON = exports.TemplateDeviceLocationsCriteriaFromJSONTyped = exports.TemplateDeviceLocationsCriteriaFromJSON = exports.instanceOfTemplateDeviceLocationsCriteria = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const TemplateDeviceLocationType_1 = require("./TemplateDeviceLocationType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the TemplateDeviceLocationsCriteria interface.
 */
function instanceOfTemplateDeviceLocationsCriteria(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfTemplateDeviceLocationsCriteria = instanceOfTemplateDeviceLocationsCriteria;
function TemplateDeviceLocationsCriteriaFromJSON(json) {
    return TemplateDeviceLocationsCriteriaFromJSONTyped(json, false);
}
exports.TemplateDeviceLocationsCriteriaFromJSON = TemplateDeviceLocationsCriteriaFromJSON;
function TemplateDeviceLocationsCriteriaFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'templateDeviceLocations': !(0, runtime_1.exists)(json, 'templateDeviceLocations') ? undefined : (json['templateDeviceLocations'].map(TemplateDeviceLocationType_1.TemplateDeviceLocationTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.TemplateDeviceLocationsCriteriaFromJSONTyped = TemplateDeviceLocationsCriteriaFromJSONTyped;
function TemplateDeviceLocationsCriteriaToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'templateDeviceLocations': value.templateDeviceLocations === undefined ? undefined : (value.templateDeviceLocations.map(TemplateDeviceLocationType_1.TemplateDeviceLocationTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.TemplateDeviceLocationsCriteriaToJSON = TemplateDeviceLocationsCriteriaToJSON;
