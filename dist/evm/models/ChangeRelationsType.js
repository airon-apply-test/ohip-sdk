"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ChangeRelationsTypeToJSON = exports.ChangeRelationsTypeFromJSONTyped = exports.ChangeRelationsTypeFromJSON = exports.instanceOfChangeRelationsType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the ChangeRelationsType interface.
 */
function instanceOfChangeRelationsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfChangeRelationsType = instanceOfChangeRelationsType;
function ChangeRelationsTypeFromJSON(json) {
    return ChangeRelationsTypeFromJSONTyped(json, false);
}
exports.ChangeRelationsTypeFromJSON = ChangeRelationsTypeFromJSON;
function ChangeRelationsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'id': !(0, runtime_1.exists)(json, 'id') ? undefined : json['id'],
        'relation': !(0, runtime_1.exists)(json, 'relation') ? undefined : json['relation'],
        'relationDescription': !(0, runtime_1.exists)(json, 'relationDescription') ? undefined : json['relationDescription'],
        'targetRelation': !(0, runtime_1.exists)(json, 'targetRelation') ? undefined : json['targetRelation'],
        'targetRelationDescription': !(0, runtime_1.exists)(json, 'targetRelationDescription') ? undefined : json['targetRelationDescription'],
    };
}
exports.ChangeRelationsTypeFromJSONTyped = ChangeRelationsTypeFromJSONTyped;
function ChangeRelationsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'id': value.id,
        'relation': value.relation,
        'relationDescription': value.relationDescription,
        'targetRelation': value.targetRelation,
        'targetRelationDescription': value.targetRelationDescription,
    };
}
exports.ChangeRelationsTypeToJSON = ChangeRelationsTypeToJSON;
