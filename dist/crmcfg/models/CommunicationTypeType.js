"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud CRM Configuration API
 * APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.CommunicationTypeTypeToJSON = exports.CommunicationTypeTypeFromJSONTyped = exports.CommunicationTypeTypeFromJSON = exports.instanceOfCommunicationTypeType = void 0;
const runtime_1 = require("../runtime");
const CommunicationRoleType_1 = require("./CommunicationRoleType");
/**
 * Check if a given object implements the CommunicationTypeType interface.
 */
function instanceOfCommunicationTypeType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCommunicationTypeType = instanceOfCommunicationTypeType;
function CommunicationTypeTypeFromJSON(json) {
    return CommunicationTypeTypeFromJSONTyped(json, false);
}
exports.CommunicationTypeTypeFromJSON = CommunicationTypeTypeFromJSON;
function CommunicationTypeTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'role': !(0, runtime_1.exists)(json, 'role') ? undefined : (0, CommunicationRoleType_1.CommunicationRoleTypeFromJSON)(json['role']),
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'textEnabled': !(0, runtime_1.exists)(json, 'textEnabled') ? undefined : json['textEnabled'],
        'sequence': !(0, runtime_1.exists)(json, 'sequence') ? undefined : json['sequence'],
        'code': !(0, runtime_1.exists)(json, 'code') ? undefined : json['code'],
    };
}
exports.CommunicationTypeTypeFromJSONTyped = CommunicationTypeTypeFromJSONTyped;
function CommunicationTypeTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'role': (0, CommunicationRoleType_1.CommunicationRoleTypeToJSON)(value.role),
        'description': value.description,
        'textEnabled': value.textEnabled,
        'sequence': value.sequence,
        'code': value.code,
    };
}
exports.CommunicationTypeTypeToJSON = CommunicationTypeTypeToJSON;
