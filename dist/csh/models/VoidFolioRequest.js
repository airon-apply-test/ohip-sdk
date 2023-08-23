"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.VoidFolioRequestToJSON = exports.VoidFolioRequestFromJSONTyped = exports.VoidFolioRequestFromJSON = exports.instanceOfVoidFolioRequest = void 0;
const runtime_1 = require("../runtime");
const InstanceLink_1 = require("./InstanceLink");
const VoidFolioCriteriaType_1 = require("./VoidFolioCriteriaType");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the VoidFolioRequest interface.
 */
function instanceOfVoidFolioRequest(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfVoidFolioRequest = instanceOfVoidFolioRequest;
function VoidFolioRequestFromJSON(json) {
    return VoidFolioRequestFromJSONTyped(json, false);
}
exports.VoidFolioRequestFromJSON = VoidFolioRequestFromJSON;
function VoidFolioRequestFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'criteria': !(0, runtime_1.exists)(json, 'criteria') ? undefined : (0, VoidFolioCriteriaType_1.VoidFolioCriteriaTypeFromJSON)(json['criteria']),
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.VoidFolioRequestFromJSONTyped = VoidFolioRequestFromJSONTyped;
function VoidFolioRequestToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'criteria': (0, VoidFolioCriteriaType_1.VoidFolioCriteriaTypeToJSON)(value.criteria),
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.VoidFolioRequestToJSON = VoidFolioRequestToJSON;
