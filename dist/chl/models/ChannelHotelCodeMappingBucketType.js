"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Channel Configuration API
 * APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ChannelHotelCodeMappingBucketTypeToJSON = exports.ChannelHotelCodeMappingBucketTypeFromJSONTyped = exports.ChannelHotelCodeMappingBucketTypeFromJSON = exports.instanceOfChannelHotelCodeMappingBucketType = void 0;
const runtime_1 = require("../runtime");
/**
 * Check if a given object implements the ChannelHotelCodeMappingBucketType interface.
 */
function instanceOfChannelHotelCodeMappingBucketType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfChannelHotelCodeMappingBucketType = instanceOfChannelHotelCodeMappingBucketType;
function ChannelHotelCodeMappingBucketTypeFromJSON(json) {
    return ChannelHotelCodeMappingBucketTypeFromJSONTyped(json, false);
}
exports.ChannelHotelCodeMappingBucketTypeFromJSON = ChannelHotelCodeMappingBucketTypeFromJSON;
function ChannelHotelCodeMappingBucketTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'lowRevenueThreshold': !(0, runtime_1.exists)(json, 'lowRevenueThreshold') ? undefined : json['lowRevenueThreshold'],
        'highRevenueThreshold': !(0, runtime_1.exists)(json, 'highRevenueThreshold') ? undefined : json['highRevenueThreshold'],
        'defaultRateCode': !(0, runtime_1.exists)(json, 'defaultRateCode') ? undefined : json['defaultRateCode'],
    };
}
exports.ChannelHotelCodeMappingBucketTypeFromJSONTyped = ChannelHotelCodeMappingBucketTypeFromJSONTyped;
function ChannelHotelCodeMappingBucketTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'lowRevenueThreshold': value.lowRevenueThreshold,
        'highRevenueThreshold': value.highRevenueThreshold,
        'defaultRateCode': value.defaultRateCode,
    };
}
exports.ChannelHotelCodeMappingBucketTypeToJSON = ChannelHotelCodeMappingBucketTypeToJSON;
