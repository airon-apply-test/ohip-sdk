"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation Asynchronous API
 * APIs to cater for Reservation Asynchronous functionality in OPERA Cloud. This includes viewing reservation data along with its revenue. <p>This API follows an async pattern where</p><ul><li>You make an initial request, which returns a Location header</li><li>You poll HEAD on the Location header returned to obtain the status of the process</li><li>Once the process completes HEAD will return in the Location header the URL that must be called to obtain the results of the process</li><li>You call the URL to obtain the results of the process</li></ul><br /><br /> Compatible with OPERA Cloud release 22.4.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.4.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.PMSResStatusTypeToJSON = exports.PMSResStatusTypeFromJSONTyped = exports.PMSResStatusTypeFromJSON = exports.PMSResStatusType = void 0;
/**
 * This reservation is in checked in status and the business date is past departure date. This could occur when ORS and PMS are in same environment.
 * @export
 */
exports.PMSResStatusType = {
    Reserved: 'Reserved',
    Requested: 'Requested',
    NoShow: 'NoShow',
    Cancelled: 'Cancelled',
    InHouse: 'InHouse',
    CheckedIn: 'CheckedIn',
    CheckedOut: 'CheckedOut',
    Waitlisted: 'Waitlisted',
    DueIn: 'DueIn',
    DueOut: 'DueOut',
    Walkin: 'Walkin',
    PendingCheckout: 'PendingCheckout'
};
function PMSResStatusTypeFromJSON(json) {
    return PMSResStatusTypeFromJSONTyped(json, false);
}
exports.PMSResStatusTypeFromJSON = PMSResStatusTypeFromJSON;
function PMSResStatusTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.PMSResStatusTypeFromJSONTyped = PMSResStatusTypeFromJSONTyped;
function PMSResStatusTypeToJSON(value) {
    return value;
}
exports.PMSResStatusTypeToJSON = PMSResStatusTypeToJSON;
