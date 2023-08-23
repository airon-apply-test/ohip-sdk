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
exports.ResPreConfiguredRoutingInstrTypeToJSON = exports.ResPreConfiguredRoutingInstrTypeFromJSONTyped = exports.ResPreConfiguredRoutingInstrTypeFromJSON = exports.instanceOfResPreConfiguredRoutingInstrType = void 0;
const runtime_1 = require("../runtime");
const AuthorizerInfoType_1 = require("./AuthorizerInfoType");
const ResProfileTypeType_1 = require("./ResProfileTypeType");
/**
 * Check if a given object implements the ResPreConfiguredRoutingInstrType interface.
 */
function instanceOfResPreConfiguredRoutingInstrType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfResPreConfiguredRoutingInstrType = instanceOfResPreConfiguredRoutingInstrType;
function ResPreConfiguredRoutingInstrTypeFromJSON(json) {
    return ResPreConfiguredRoutingInstrTypeFromJSONTyped(json, false);
}
exports.ResPreConfiguredRoutingInstrTypeFromJSON = ResPreConfiguredRoutingInstrTypeFromJSON;
function ResPreConfiguredRoutingInstrTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'authorizerInfo': !(0, runtime_1.exists)(json, 'authorizerInfo') ? undefined : (0, AuthorizerInfoType_1.AuthorizerInfoTypeFromJSON)(json['authorizerInfo']),
        'ratePlanCode': !(0, runtime_1.exists)(json, 'ratePlanCode') ? undefined : json['ratePlanCode'],
        'profileType': !(0, runtime_1.exists)(json, 'profileType') ? undefined : (0, ResProfileTypeType_1.ResProfileTypeTypeFromJSON)(json['profileType']),
        'promotionCode': !(0, runtime_1.exists)(json, 'promotionCode') ? undefined : json['promotionCode'],
    };
}
exports.ResPreConfiguredRoutingInstrTypeFromJSONTyped = ResPreConfiguredRoutingInstrTypeFromJSONTyped;
function ResPreConfiguredRoutingInstrTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'authorizerInfo': (0, AuthorizerInfoType_1.AuthorizerInfoTypeToJSON)(value.authorizerInfo),
        'ratePlanCode': value.ratePlanCode,
        'profileType': (0, ResProfileTypeType_1.ResProfileTypeTypeToJSON)(value.profileType),
        'promotionCode': value.promotionCode,
    };
}
exports.ResPreConfiguredRoutingInstrTypeToJSON = ResPreConfiguredRoutingInstrTypeToJSON;
