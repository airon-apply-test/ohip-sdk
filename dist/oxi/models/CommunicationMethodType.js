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
exports.CommunicationMethodTypeToJSON = exports.CommunicationMethodTypeFromJSONTyped = exports.CommunicationMethodTypeFromJSON = exports.instanceOfCommunicationMethodType = void 0;
const runtime_1 = require("../runtime");
const CommunicationMethodHTTPSType_1 = require("./CommunicationMethodHTTPSType");
const CommunicationMethodNoneType_1 = require("./CommunicationMethodNoneType");
const CommunicationType_1 = require("./CommunicationType");
/**
 * Check if a given object implements the CommunicationMethodType interface.
 */
function instanceOfCommunicationMethodType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCommunicationMethodType = instanceOfCommunicationMethodType;
function CommunicationMethodTypeFromJSON(json) {
    return CommunicationMethodTypeFromJSONTyped(json, false);
}
exports.CommunicationMethodTypeFromJSON = CommunicationMethodTypeFromJSON;
function CommunicationMethodTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'interfaceId': !(0, runtime_1.exists)(json, 'interfaceId') ? undefined : json['interfaceId'],
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'communicationType': !(0, runtime_1.exists)(json, 'communicationType') ? undefined : (0, CommunicationType_1.CommunicationTypeFromJSON)(json['communicationType']),
        'communicationFlow': !(0, runtime_1.exists)(json, 'communicationFlow') ? undefined : json['communicationFlow'],
        'noCommunication': !(0, runtime_1.exists)(json, 'noCommunication') ? undefined : (0, CommunicationMethodNoneType_1.CommunicationMethodNoneTypeFromJSON)(json['noCommunication']),
        'hTTPS': !(0, runtime_1.exists)(json, 'hTTPS') ? undefined : (0, CommunicationMethodHTTPSType_1.CommunicationMethodHTTPSTypeFromJSON)(json['hTTPS']),
        'allowedComTypes': !(0, runtime_1.exists)(json, 'allowedComTypes') ? undefined : json['allowedComTypes'],
        'displayText': !(0, runtime_1.exists)(json, 'displayText') ? undefined : json['displayText'],
        'retryCount': !(0, runtime_1.exists)(json, 'retryCount') ? undefined : json['retryCount'],
        'retryInterval': !(0, runtime_1.exists)(json, 'retryInterval') ? undefined : json['retryInterval'],
        'startTime': !(0, runtime_1.exists)(json, 'startTime') ? undefined : json['startTime'],
        'noOfDaysToKeep': !(0, runtime_1.exists)(json, 'noOfDaysToKeep') ? undefined : json['noOfDaysToKeep'],
        'versionNumber': !(0, runtime_1.exists)(json, 'versionNumber') ? undefined : json['versionNumber'],
        'systemType': !(0, runtime_1.exists)(json, 'systemType') ? undefined : json['systemType'],
        'chainCode': !(0, runtime_1.exists)(json, 'chainCode') ? undefined : json['chainCode'],
    };
}
exports.CommunicationMethodTypeFromJSONTyped = CommunicationMethodTypeFromJSONTyped;
function CommunicationMethodTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'interfaceId': value.interfaceId,
        'hotelId': value.hotelId,
        'communicationType': (0, CommunicationType_1.CommunicationTypeToJSON)(value.communicationType),
        'communicationFlow': value.communicationFlow,
        'noCommunication': (0, CommunicationMethodNoneType_1.CommunicationMethodNoneTypeToJSON)(value.noCommunication),
        'hTTPS': (0, CommunicationMethodHTTPSType_1.CommunicationMethodHTTPSTypeToJSON)(value.hTTPS),
        'allowedComTypes': value.allowedComTypes,
        'displayText': value.displayText,
        'retryCount': value.retryCount,
        'retryInterval': value.retryInterval,
        'startTime': value.startTime,
        'noOfDaysToKeep': value.noOfDaysToKeep,
        'versionNumber': value.versionNumber,
        'systemType': value.systemType,
        'chainCode': value.chainCode,
    };
}
exports.CommunicationMethodTypeToJSON = CommunicationMethodTypeToJSON;
