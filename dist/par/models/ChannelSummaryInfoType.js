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
exports.ChannelSummaryInfoTypeToJSON = exports.ChannelSummaryInfoTypeFromJSONTyped = exports.ChannelSummaryInfoTypeFromJSON = exports.instanceOfChannelSummaryInfoType = void 0;
const runtime_1 = require("../runtime");
const BookingChannelType_1 = require("./BookingChannelType");
/**
 * Check if a given object implements the ChannelSummaryInfoType interface.
 */
function instanceOfChannelSummaryInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfChannelSummaryInfoType = instanceOfChannelSummaryInfoType;
function ChannelSummaryInfoTypeFromJSON(json) {
    return ChannelSummaryInfoTypeFromJSONTyped(json, false);
}
exports.ChannelSummaryInfoTypeFromJSON = ChannelSummaryInfoTypeFromJSON;
function ChannelSummaryInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'bookingChannel': !(0, runtime_1.exists)(json, 'bookingChannel') ? undefined : (0, BookingChannelType_1.BookingChannelTypeFromJSON)(json['bookingChannel']),
        'messageId': !(0, runtime_1.exists)(json, 'messageId') ? undefined : json['messageId'],
        'summaryOnly': !(0, runtime_1.exists)(json, 'summaryOnly') ? undefined : json['summaryOnly'],
    };
}
exports.ChannelSummaryInfoTypeFromJSONTyped = ChannelSummaryInfoTypeFromJSONTyped;
function ChannelSummaryInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'bookingChannel': (0, BookingChannelType_1.BookingChannelTypeToJSON)(value.bookingChannel),
        'messageId': value.messageId,
        'summaryOnly': value.summaryOnly,
    };
}
exports.ChannelSummaryInfoTypeToJSON = ChannelSummaryInfoTypeToJSON;
