"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Activity API
 * APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.RelationshipProfileTypeToJSON = exports.RelationshipProfileTypeFromJSONTyped = exports.RelationshipProfileTypeFromJSON = exports.instanceOfRelationshipProfileType = void 0;
const runtime_1 = require("../runtime");
const ProfileTypeType_1 = require("./ProfileTypeType");
const UniqueIDType_1 = require("./UniqueIDType");
/**
 * Check if a given object implements the RelationshipProfileType interface.
 */
function instanceOfRelationshipProfileType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfRelationshipProfileType = instanceOfRelationshipProfileType;
function RelationshipProfileTypeFromJSON(json) {
    return RelationshipProfileTypeFromJSONTyped(json, false);
}
exports.RelationshipProfileTypeFromJSON = RelationshipProfileTypeFromJSON;
function RelationshipProfileTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'profileIdList': !(0, runtime_1.exists)(json, 'profileIdList') ? undefined : (json['profileIdList'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'profileType': !(0, runtime_1.exists)(json, 'profileType') ? undefined : (0, ProfileTypeType_1.ProfileTypeTypeFromJSON)(json['profileType']),
    };
}
exports.RelationshipProfileTypeFromJSONTyped = RelationshipProfileTypeFromJSONTyped;
function RelationshipProfileTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'profileIdList': value.profileIdList === undefined ? undefined : (value.profileIdList.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'profileType': (0, ProfileTypeType_1.ProfileTypeTypeToJSON)(value.profileType),
    };
}
exports.RelationshipProfileTypeToJSON = RelationshipProfileTypeToJSON;
