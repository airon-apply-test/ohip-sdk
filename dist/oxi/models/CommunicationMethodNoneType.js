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
exports.CommunicationMethodNoneTypeToJSON = exports.CommunicationMethodNoneTypeFromJSONTyped = exports.CommunicationMethodNoneTypeFromJSON = exports.instanceOfCommunicationMethodNoneType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the CommunicationMethodNoneType interface.
 */
function instanceOfCommunicationMethodNoneType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCommunicationMethodNoneType = instanceOfCommunicationMethodNoneType;
function CommunicationMethodNoneTypeFromJSON(json) {
    return CommunicationMethodNoneTypeFromJSONTyped(json, false);
}
exports.CommunicationMethodNoneTypeFromJSON = CommunicationMethodNoneTypeFromJSON;
function CommunicationMethodNoneTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'sleepTime': !(0, runtime_1.exists)(json, 'sleepTime') ? undefined : json['sleepTime'],
    };
}
exports.CommunicationMethodNoneTypeFromJSONTyped = CommunicationMethodNoneTypeFromJSONTyped;
function CommunicationMethodNoneTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'sleepTime': value.sleepTime,
    };
}
exports.CommunicationMethodNoneTypeToJSON = CommunicationMethodNoneTypeToJSON;
