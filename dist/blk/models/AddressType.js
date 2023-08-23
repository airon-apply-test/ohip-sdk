"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.AddressTypeToJSON = exports.AddressTypeFromJSONTyped = exports.AddressTypeFromJSON = exports.instanceOfAddressType = void 0;
const runtime_1 = require("../runtime");
const CountryNameType_1 = require("./CountryNameType");
/**
 * Check if a given object implements the AddressType interface.
 */
function instanceOfAddressType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfAddressType = instanceOfAddressType;
function AddressTypeFromJSON(json) {
    return AddressTypeFromJSONTyped(json, false);
}
exports.AddressTypeFromJSON = AddressTypeFromJSON;
function AddressTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'isValidated': !(0, runtime_1.exists)(json, 'isValidated') ? undefined : json['isValidated'],
        'addressLine': !(0, runtime_1.exists)(json, 'addressLine') ? undefined : json['addressLine'],
        'cityName': !(0, runtime_1.exists)(json, 'cityName') ? undefined : json['cityName'],
        'postalCode': !(0, runtime_1.exists)(json, 'postalCode') ? undefined : json['postalCode'],
        'cityExtension': !(0, runtime_1.exists)(json, 'cityExtension') ? undefined : json['cityExtension'],
        'county': !(0, runtime_1.exists)(json, 'county') ? undefined : json['county'],
        'state': !(0, runtime_1.exists)(json, 'state') ? undefined : json['state'],
        'country': !(0, runtime_1.exists)(json, 'country') ? undefined : (0, CountryNameType_1.CountryNameTypeFromJSON)(json['country']),
        'language': !(0, runtime_1.exists)(json, 'language') ? undefined : json['language'],
        'type': !(0, runtime_1.exists)(json, 'type') ? undefined : json['type'],
        'typeDescription': !(0, runtime_1.exists)(json, 'typeDescription') ? undefined : json['typeDescription'],
        'primaryInd': !(0, runtime_1.exists)(json, 'primaryInd') ? undefined : json['primaryInd'],
        'updateReservations': !(0, runtime_1.exists)(json, 'updateReservations') ? undefined : json['updateReservations'],
        'createDateTime': !(0, runtime_1.exists)(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !(0, runtime_1.exists)(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !(0, runtime_1.exists)(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !(0, runtime_1.exists)(json, 'lastModifierId') ? undefined : json['lastModifierId'],
    };
}
exports.AddressTypeFromJSONTyped = AddressTypeFromJSONTyped;
function AddressTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'isValidated': value.isValidated,
        'addressLine': value.addressLine,
        'cityName': value.cityName,
        'postalCode': value.postalCode,
        'cityExtension': value.cityExtension,
        'county': value.county,
        'state': value.state,
        'country': (0, CountryNameType_1.CountryNameTypeToJSON)(value.country),
        'language': value.language,
        'type': value.type,
        'typeDescription': value.typeDescription,
        'primaryInd': value.primaryInd,
        'updateReservations': value.updateReservations,
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
    };
}
exports.AddressTypeToJSON = AddressTypeToJSON;
