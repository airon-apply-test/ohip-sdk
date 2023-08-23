"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Content Service
 * Opera Content Service offers capability to manage large content such as images and files.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.RegistrationCardTypeToJSON = exports.RegistrationCardTypeFromJSONTyped = exports.RegistrationCardTypeFromJSON = exports.instanceOfRegistrationCardType = void 0;
const runtime_1 = require("../runtime");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the RegistrationCardType interface.
 */
function instanceOfRegistrationCardType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRegistrationCardType = instanceOfRegistrationCardType;
function RegistrationCardTypeFromJSON(json) {
    return RegistrationCardTypeFromJSONTyped(json, false);
}
exports.RegistrationCardTypeFromJSON = RegistrationCardTypeFromJSON;
function RegistrationCardTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'reservationId': !(0, runtime_1.exists)(json, 'reservationId') ? undefined : (0, UniqueIDType_1.UniqueIDTypeFromJSON)(json['reservationId']),
        'registrationCard': !(0, runtime_1.exists)(json, 'registrationCard') ? undefined : json['registrationCard'],
    };
}
exports.RegistrationCardTypeFromJSONTyped = RegistrationCardTypeFromJSONTyped;
function RegistrationCardTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'hotelId': value.hotelId,
        'reservationId': (0, UniqueIDType_1.UniqueIDTypeToJSON)(value.reservationId),
        'registrationCard': value.registrationCard,
    };
}
exports.RegistrationCardTypeToJSON = RegistrationCardTypeToJSON;
