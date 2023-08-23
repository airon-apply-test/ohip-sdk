"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud API for Customer Management Service
 * This API deals with the different aspect of the CustomerManagement.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.PersonNameTypeToJSON = exports.PersonNameTypeFromJSONTyped = exports.PersonNameTypeFromJSON = exports.instanceOfPersonNameType = void 0;
const runtime_1 = require("../runtime");
const PersonNameTypeType_1 = require("./PersonNameTypeType");
/**
 * Check if a given object implements the PersonNameType interface.
 */
function instanceOfPersonNameType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfPersonNameType = instanceOfPersonNameType;
function PersonNameTypeFromJSON(json) {
    return PersonNameTypeFromJSONTyped(json, false);
}
exports.PersonNameTypeFromJSON = PersonNameTypeFromJSON;
function PersonNameTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'namePrefix': !(0, runtime_1.exists)(json, 'namePrefix') ? undefined : json['namePrefix'],
        'givenName': !(0, runtime_1.exists)(json, 'givenName') ? undefined : json['givenName'],
        'middleName': !(0, runtime_1.exists)(json, 'middleName') ? undefined : json['middleName'],
        'surname': !(0, runtime_1.exists)(json, 'surname') ? undefined : json['surname'],
        'nameSuffix': !(0, runtime_1.exists)(json, 'nameSuffix') ? undefined : json['nameSuffix'],
        'nameTitle': !(0, runtime_1.exists)(json, 'nameTitle') ? undefined : json['nameTitle'],
        'nameTitleSuffix': !(0, runtime_1.exists)(json, 'nameTitleSuffix') ? undefined : json['nameTitleSuffix'],
        'envelopeGreeting': !(0, runtime_1.exists)(json, 'envelopeGreeting') ? undefined : json['envelopeGreeting'],
        'salutation': !(0, runtime_1.exists)(json, 'salutation') ? undefined : json['salutation'],
        'nameType': !(0, runtime_1.exists)(json, 'nameType') ? undefined : (0, PersonNameTypeType_1.PersonNameTypeTypeFromJSON)(json['nameType']),
        'language': !(0, runtime_1.exists)(json, 'language') ? undefined : json['language'],
        'externalSystem': !(0, runtime_1.exists)(json, 'externalSystem') ? undefined : json['externalSystem'],
    };
}
exports.PersonNameTypeFromJSONTyped = PersonNameTypeFromJSONTyped;
function PersonNameTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'namePrefix': value.namePrefix,
        'givenName': value.givenName,
        'middleName': value.middleName,
        'surname': value.surname,
        'nameSuffix': value.nameSuffix,
        'nameTitle': value.nameTitle,
        'nameTitleSuffix': value.nameTitleSuffix,
        'envelopeGreeting': value.envelopeGreeting,
        'salutation': value.salutation,
        'nameType': (0, PersonNameTypeType_1.PersonNameTypeTypeToJSON)(value.nameType),
        'language': value.language,
        'externalSystem': value.externalSystem,
    };
}
exports.PersonNameTypeToJSON = PersonNameTypeToJSON;
