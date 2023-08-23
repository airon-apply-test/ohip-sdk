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
exports.HotelInfoTypeAddressToJSON = exports.HotelInfoTypeAddressFromJSONTyped = exports.HotelInfoTypeAddressFromJSON = exports.instanceOfHotelInfoTypeAddress = void 0;
const runtime_1 = require("../runtime");
const CountryNameType_1 = require("./CountryNameType");
/**
 * Check if a given object implements the HotelInfoTypeAddress interface.
 */
function instanceOfHotelInfoTypeAddress(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfHotelInfoTypeAddress = instanceOfHotelInfoTypeAddress;
function HotelInfoTypeAddressFromJSON(json) {
    return HotelInfoTypeAddressFromJSONTyped(json, false);
}
exports.HotelInfoTypeAddressFromJSON = HotelInfoTypeAddressFromJSON;
function HotelInfoTypeAddressFromJSONTyped(json, ignoreDiscriminator) {
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
        'barCode': !(0, runtime_1.exists)(json, 'barCode') ? undefined : json['barCode'],
        'createDateTime': !(0, runtime_1.exists)(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !(0, runtime_1.exists)(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !(0, runtime_1.exists)(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !(0, runtime_1.exists)(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'purgeDate': !(0, runtime_1.exists)(json, 'purgeDate') ? undefined : (new Date(json['purgeDate'])),
        'regionCode': !(0, runtime_1.exists)(json, 'regionCode') ? undefined : json['regionCode'],
    };
}
exports.HotelInfoTypeAddressFromJSONTyped = HotelInfoTypeAddressFromJSONTyped;
function HotelInfoTypeAddressToJSON(value) {
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
        'barCode': value.barCode,
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'purgeDate': value.purgeDate === undefined ? undefined : (value.purgeDate.toISOString().substr(0, 10)),
        'regionCode': value.regionCode,
    };
}
exports.HotelInfoTypeAddressToJSON = HotelInfoTypeAddressToJSON;
