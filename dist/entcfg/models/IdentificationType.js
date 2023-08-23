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
exports.IdentificationTypeToJSON = exports.IdentificationTypeFromJSONTyped = exports.IdentificationTypeFromJSON = exports.instanceOfIdentificationType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the IdentificationType interface.
 */
function instanceOfIdentificationType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfIdentificationType = instanceOfIdentificationType;
function IdentificationTypeFromJSON(json) {
    return IdentificationTypeFromJSONTyped(json, false);
}
exports.IdentificationTypeFromJSON = IdentificationTypeFromJSON;
function IdentificationTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'idType': !(0, runtime_1.exists)(json, 'idType') ? undefined : json['idType'],
        'idNumber': !(0, runtime_1.exists)(json, 'idNumber') ? undefined : json['idNumber'],
        'idNumberMasked': !(0, runtime_1.exists)(json, 'idNumberMasked') ? undefined : json['idNumberMasked'],
        'issuedCountry': !(0, runtime_1.exists)(json, 'issuedCountry') ? undefined : json['issuedCountry'],
        'issuedDate': !(0, runtime_1.exists)(json, 'issuedDate') ? undefined : (new Date(json['issuedDate'])),
        'issuedPlace': !(0, runtime_1.exists)(json, 'issuedPlace') ? undefined : json['issuedPlace'],
        'expirationDate': !(0, runtime_1.exists)(json, 'expirationDate') ? undefined : (new Date(json['expirationDate'])),
        'registeredProperty': !(0, runtime_1.exists)(json, 'registeredProperty') ? undefined : json['registeredProperty'],
        'primaryInd': !(0, runtime_1.exists)(json, 'primaryInd') ? undefined : json['primaryInd'],
        'orderSequence': !(0, runtime_1.exists)(json, 'orderSequence') ? undefined : json['orderSequence'],
    };
}
exports.IdentificationTypeFromJSONTyped = IdentificationTypeFromJSONTyped;
function IdentificationTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'idType': value.idType,
        'idNumber': value.idNumber,
        'idNumberMasked': value.idNumberMasked,
        'issuedCountry': value.issuedCountry,
        'issuedDate': value.issuedDate === undefined ? undefined : (value.issuedDate.toISOString().substr(0, 10)),
        'issuedPlace': value.issuedPlace,
        'expirationDate': value.expirationDate === undefined ? undefined : (value.expirationDate.toISOString().substr(0, 10)),
        'registeredProperty': value.registeredProperty,
        'primaryInd': value.primaryInd,
        'orderSequence': value.orderSequence,
    };
}
exports.IdentificationTypeToJSON = IdentificationTypeToJSON;
