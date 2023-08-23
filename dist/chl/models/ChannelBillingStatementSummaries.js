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
exports.ChannelBillingStatementSummariesToJSON = exports.ChannelBillingStatementSummariesFromJSONTyped = exports.ChannelBillingStatementSummariesFromJSON = exports.instanceOfChannelBillingStatementSummaries = void 0;
const runtime_1 = require("../runtime");
const ChannelBillingStatementSummaryType_1 = require("./ChannelBillingStatementSummaryType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the ChannelBillingStatementSummaries interface.
 */
function instanceOfChannelBillingStatementSummaries(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfChannelBillingStatementSummaries = instanceOfChannelBillingStatementSummaries;
function ChannelBillingStatementSummariesFromJSON(json) {
    return ChannelBillingStatementSummariesFromJSONTyped(json, false);
}
exports.ChannelBillingStatementSummariesFromJSON = ChannelBillingStatementSummariesFromJSON;
function ChannelBillingStatementSummariesFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'channelBillingStatementsSummary': !(0, runtime_1.exists)(json, 'channelBillingStatementsSummary') ? undefined : (json['channelBillingStatementsSummary'].map(ChannelBillingStatementSummaryType_1.ChannelBillingStatementSummaryTypeFromJSON)),
        'totalPages': !(0, runtime_1.exists)(json, 'totalPages') ? undefined : json['totalPages'],
        'offset': !(0, runtime_1.exists)(json, 'offset') ? undefined : json['offset'],
        'limit': !(0, runtime_1.exists)(json, 'limit') ? undefined : json['limit'],
        'hasMore': !(0, runtime_1.exists)(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !(0, runtime_1.exists)(json, 'totalResults') ? undefined : json['totalResults'],
        'count': !(0, runtime_1.exists)(json, 'count') ? undefined : json['count'],
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.ChannelBillingStatementSummariesFromJSONTyped = ChannelBillingStatementSummariesFromJSONTyped;
function ChannelBillingStatementSummariesToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'channelBillingStatementsSummary': value.channelBillingStatementsSummary === undefined ? undefined : (value.channelBillingStatementsSummary.map(ChannelBillingStatementSummaryType_1.ChannelBillingStatementSummaryTypeToJSON)),
        'totalPages': value.totalPages,
        'offset': value.offset,
        'limit': value.limit,
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'count': value.count,
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.ChannelBillingStatementSummariesToJSON = ChannelBillingStatementSummariesToJSON;
