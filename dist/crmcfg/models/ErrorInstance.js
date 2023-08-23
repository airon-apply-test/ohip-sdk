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
exports.ErrorInstanceToJSON = exports.ErrorInstanceFromJSONTyped = exports.ErrorInstanceFromJSON = exports.instanceOfErrorInstance = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the ErrorInstance interface.
 */
function instanceOfErrorInstance(value) {
    let isInstance = true;
    isInstance = isInstance && "type" in value;
    isInstance = isInstance && "title" in value;
    return isInstance;
}
exports.instanceOfErrorInstance = instanceOfErrorInstance;
function ErrorInstanceFromJSON(json) {
    return ErrorInstanceFromJSONTyped(json, false);
}
exports.ErrorInstanceFromJSON = ErrorInstanceFromJSON;
function ErrorInstanceFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'type': json['type'],
        'title': json['title'],
        'status': !(0, runtime_1.exists)(json, 'status') ? undefined : json['status'],
        'detail': !(0, runtime_1.exists)(json, 'detail') ? undefined : json['detail'],
        'instance': !(0, runtime_1.exists)(json, 'instance') ? undefined : json['instance'],
        'oerrorCode': !(0, runtime_1.exists)(json, 'o:errorCode') ? undefined : json['o:errorCode'],
        'oerrorPath': !(0, runtime_1.exists)(json, 'o:errorPath') ? undefined : json['o:errorPath'],
    };
}
exports.ErrorInstanceFromJSONTyped = ErrorInstanceFromJSONTyped;
function ErrorInstanceToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'type': value.type,
        'title': value.title,
        'status': value.status,
        'detail': value.detail,
        'instance': value.instance,
        'o:errorCode': value.oerrorCode,
        'o:errorPath': value.oerrorPath,
    };
}
exports.ErrorInstanceToJSON = ErrorInstanceToJSON;
