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
exports.NegotiatedInfoTypeToJSON = exports.NegotiatedInfoTypeFromJSONTyped = exports.NegotiatedInfoTypeFromJSON = exports.instanceOfNegotiatedInfoType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the NegotiatedInfoType interface.
 */
function instanceOfNegotiatedInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfNegotiatedInfoType = instanceOfNegotiatedInfoType;
function NegotiatedInfoTypeFromJSON(json) {
    return NegotiatedInfoTypeFromJSONTyped(json, false);
}
exports.NegotiatedInfoTypeFromJSON = NegotiatedInfoTypeFromJSON;
function NegotiatedInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'corporateAgreementId': !(0, runtime_1.exists)(json, 'corporateAgreementId') ? undefined : json['corporateAgreementId'],
        'comissionCode': !(0, runtime_1.exists)(json, 'comissionCode') ? undefined : json['comissionCode'],
        'order': !(0, runtime_1.exists)(json, 'order') ? undefined : json['order'],
        'inactive': !(0, runtime_1.exists)(json, 'inactive') ? undefined : json['inactive'],
        'start': !(0, runtime_1.exists)(json, 'start') ? undefined : (new Date(json['start'])),
        'end': !(0, runtime_1.exists)(json, 'end') ? undefined : (new Date(json['end'])),
    };
}
exports.NegotiatedInfoTypeFromJSONTyped = NegotiatedInfoTypeFromJSONTyped;
function NegotiatedInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'corporateAgreementId': value.corporateAgreementId,
        'comissionCode': value.comissionCode,
        'order': value.order,
        'inactive': value.inactive,
        'start': value.start === undefined ? undefined : (value.start.toISOString().substr(0, 10)),
        'end': value.end === undefined ? undefined : (value.end.toISOString().substr(0, 10)),
    };
}
exports.NegotiatedInfoTypeToJSON = NegotiatedInfoTypeToJSON;
