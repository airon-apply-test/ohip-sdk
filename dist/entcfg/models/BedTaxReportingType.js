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
exports.BedTaxReportingTypeToJSON = exports.BedTaxReportingTypeFromJSONTyped = exports.BedTaxReportingTypeFromJSON = exports.instanceOfBedTaxReportingType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the BedTaxReportingType interface.
 */
function instanceOfBedTaxReportingType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfBedTaxReportingType = instanceOfBedTaxReportingType;
function BedTaxReportingTypeFromJSON(json) {
    return BedTaxReportingTypeFromJSONTyped(json, false);
}
exports.BedTaxReportingTypeFromJSON = BedTaxReportingTypeFromJSON;
function BedTaxReportingTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'taxRegistrationNo': !(0, runtime_1.exists)(json, 'taxRegistrationNo') ? undefined : json['taxRegistrationNo'],
        'visaNumber': !(0, runtime_1.exists)(json, 'visaNumber') ? undefined : json['visaNumber'],
        'visaIssueDate': !(0, runtime_1.exists)(json, 'visaIssueDate') ? undefined : (new Date(json['visaIssueDate'])),
        'visaExpiryDate': !(0, runtime_1.exists)(json, 'visaExpiryDate') ? undefined : (new Date(json['visaExpiryDate'])),
        'taxableDays': !(0, runtime_1.exists)(json, 'taxableDays') ? undefined : json['taxableDays'],
    };
}
exports.BedTaxReportingTypeFromJSONTyped = BedTaxReportingTypeFromJSONTyped;
function BedTaxReportingTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'taxRegistrationNo': value.taxRegistrationNo,
        'visaNumber': value.visaNumber,
        'visaIssueDate': value.visaIssueDate === undefined ? undefined : (value.visaIssueDate.toISOString().substr(0, 10)),
        'visaExpiryDate': value.visaExpiryDate === undefined ? undefined : (value.visaExpiryDate.toISOString().substr(0, 10)),
        'taxableDays': value.taxableDays,
    };
}
exports.BedTaxReportingTypeToJSON = BedTaxReportingTypeToJSON;
