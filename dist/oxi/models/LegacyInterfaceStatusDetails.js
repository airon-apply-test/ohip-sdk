"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Xchange Interface OXI API
 * APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 23.0.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.LegacyInterfaceStatusDetailsToJSON = exports.LegacyInterfaceStatusDetailsFromJSONTyped = exports.LegacyInterfaceStatusDetailsFromJSON = exports.instanceOfLegacyInterfaceStatusDetails = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const InterfaceStatusType_1 = require("./InterfaceStatusType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the LegacyInterfaceStatusDetails interface.
 */
function instanceOfLegacyInterfaceStatusDetails(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfLegacyInterfaceStatusDetails = instanceOfLegacyInterfaceStatusDetails;
function LegacyInterfaceStatusDetailsFromJSON(json) {
    return LegacyInterfaceStatusDetailsFromJSONTyped(json, false);
}
exports.LegacyInterfaceStatusDetailsFromJSON = LegacyInterfaceStatusDetailsFromJSON;
function LegacyInterfaceStatusDetailsFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'interfacesStatus': !(0, runtime_1.exists)(json, 'interfacesStatus') ? undefined : (json['interfacesStatus'].map(InterfaceStatusType_1.InterfaceStatusTypeFromJSON)),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.LegacyInterfaceStatusDetailsFromJSONTyped = LegacyInterfaceStatusDetailsFromJSONTyped;
function LegacyInterfaceStatusDetailsToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'interfacesStatus': value.interfacesStatus === undefined ? undefined : (value.interfacesStatus.map(InterfaceStatusType_1.InterfaceStatusTypeToJSON)),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.LegacyInterfaceStatusDetailsToJSON = LegacyInterfaceStatusDetailsToJSON;
