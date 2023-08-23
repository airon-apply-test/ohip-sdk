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
exports.AddressSearchTypeToJSON = exports.AddressSearchTypeFromJSONTyped = exports.AddressSearchTypeFromJSON = exports.instanceOfAddressSearchType = void 0;
const runtime_1 = require("../runtime");
const CountryNameType_1 = require("./CountryNameType");
/**
 * Check if a given object implements the AddressSearchType interface.
 */
function instanceOfAddressSearchType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfAddressSearchType = instanceOfAddressSearchType;
function AddressSearchTypeFromJSON(json) {
    return AddressSearchTypeFromJSONTyped(json, false);
}
exports.AddressSearchTypeFromJSON = AddressSearchTypeFromJSON;
function AddressSearchTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'cityName': !(0, runtime_1.exists)(json, 'cityName') ? undefined : json['cityName'],
        'postalCode': !(0, runtime_1.exists)(json, 'postalCode') ? undefined : json['postalCode'],
        'state': !(0, runtime_1.exists)(json, 'state') ? undefined : json['state'],
        'country': !(0, runtime_1.exists)(json, 'country') ? undefined : (0, CountryNameType_1.CountryNameTypeFromJSON)(json['country']),
        'streetAddress': !(0, runtime_1.exists)(json, 'streetAddress') ? undefined : json['streetAddress'],
        'excludeNoCity': !(0, runtime_1.exists)(json, 'excludeNoCity') ? undefined : json['excludeNoCity'],
    };
}
exports.AddressSearchTypeFromJSONTyped = AddressSearchTypeFromJSONTyped;
function AddressSearchTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'cityName': value.cityName,
        'postalCode': value.postalCode,
        'state': value.state,
        'country': (0, CountryNameType_1.CountryNameTypeToJSON)(value.country),
        'streetAddress': value.streetAddress,
        'excludeNoCity': value.excludeNoCity,
    };
}
exports.AddressSearchTypeToJSON = AddressSearchTypeToJSON;
