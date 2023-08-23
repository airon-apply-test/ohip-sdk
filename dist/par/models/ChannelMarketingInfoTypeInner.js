"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ChannelMarketingInfoTypeInnerToJSON = exports.ChannelMarketingInfoTypeInnerFromJSONTyped = exports.ChannelMarketingInfoTypeInnerFromJSON = exports.instanceOfChannelMarketingInfoTypeInner = void 0;
const runtime_1 = require("../runtime");
const MarketingPolicyType_1 = require("./MarketingPolicyType");
/**
 * Check if a given object implements the ChannelMarketingInfoTypeInner interface.
 */
function instanceOfChannelMarketingInfoTypeInner(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfChannelMarketingInfoTypeInner = instanceOfChannelMarketingInfoTypeInner;
function ChannelMarketingInfoTypeInnerFromJSON(json) {
    return ChannelMarketingInfoTypeInnerFromJSONTyped(json, false);
}
exports.ChannelMarketingInfoTypeInnerFromJSON = ChannelMarketingInfoTypeInnerFromJSON;
function ChannelMarketingInfoTypeInnerFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'description': !(0, runtime_1.exists)(json, 'description') ? undefined : json['description'],
        'policyType': !(0, runtime_1.exists)(json, 'policyType') ? undefined : (0, MarketingPolicyType_1.MarketingPolicyTypeFromJSON)(json['policyType']),
        'ratePlanCode': !(0, runtime_1.exists)(json, 'ratePlanCode') ? undefined : json['ratePlanCode'],
    };
}
exports.ChannelMarketingInfoTypeInnerFromJSONTyped = ChannelMarketingInfoTypeInnerFromJSONTyped;
function ChannelMarketingInfoTypeInnerToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'description': value.description,
        'policyType': (0, MarketingPolicyType_1.MarketingPolicyTypeToJSON)(value.policyType),
        'ratePlanCode': value.ratePlanCode,
    };
}
exports.ChannelMarketingInfoTypeInnerToJSON = ChannelMarketingInfoTypeInnerToJSON;
