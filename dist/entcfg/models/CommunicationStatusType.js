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
exports.CommunicationStatusTypeToJSON = exports.CommunicationStatusTypeFromJSONTyped = exports.CommunicationStatusTypeFromJSON = exports.CommunicationStatusType = void 0;
/**
 * Enum to denote the Status of Readiness messages sent to Guest Devices.
 * @export
 */
exports.CommunicationStatusType = {
    Pending: 'Pending',
    Completed: 'Completed',
    Failed: 'Failed',
    Sent: 'Sent',
    Received: 'Received',
    Cancelled: 'Cancelled',
    PendingAvailability: 'PendingAvailability'
};
function CommunicationStatusTypeFromJSON(json) {
    return CommunicationStatusTypeFromJSONTyped(json, false);
}
exports.CommunicationStatusTypeFromJSON = CommunicationStatusTypeFromJSON;
function CommunicationStatusTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.CommunicationStatusTypeFromJSONTyped = CommunicationStatusTypeFromJSONTyped;
function CommunicationStatusTypeToJSON(value) {
    return value;
}
exports.CommunicationStatusTypeToJSON = CommunicationStatusTypeToJSON;
