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
exports.ECertificateIssueSourceTypeToJSON = exports.ECertificateIssueSourceTypeFromJSONTyped = exports.ECertificateIssueSourceTypeFromJSON = exports.ECertificateIssueSourceType = void 0;
/**
 * Indicates that OPERA E-Certificate is issued by external system.
 * @export
 */
exports.ECertificateIssueSourceType = {
    Opera: 'Opera',
    Web: 'Web',
    Interface: 'Interface'
};
function ECertificateIssueSourceTypeFromJSON(json) {
    return ECertificateIssueSourceTypeFromJSONTyped(json, false);
}
exports.ECertificateIssueSourceTypeFromJSON = ECertificateIssueSourceTypeFromJSON;
function ECertificateIssueSourceTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.ECertificateIssueSourceTypeFromJSONTyped = ECertificateIssueSourceTypeFromJSONTyped;
function ECertificateIssueSourceTypeToJSON(value) {
    return value;
}
exports.ECertificateIssueSourceTypeToJSON = ECertificateIssueSourceTypeToJSON;
