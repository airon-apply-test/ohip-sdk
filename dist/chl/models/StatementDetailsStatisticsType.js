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
exports.StatementDetailsStatisticsTypeToJSON = exports.StatementDetailsStatisticsTypeFromJSONTyped = exports.StatementDetailsStatisticsTypeFromJSON = exports.instanceOfStatementDetailsStatisticsType = void 0;
const runtime_1 = require("../runtime");
const StatementDetailsStatisticType_1 = require("./StatementDetailsStatisticType");
/**
 * Check if a given object implements the StatementDetailsStatisticsType interface.
 */
function instanceOfStatementDetailsStatisticsType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfStatementDetailsStatisticsType = instanceOfStatementDetailsStatisticsType;
function StatementDetailsStatisticsTypeFromJSON(json) {
    return StatementDetailsStatisticsTypeFromJSONTyped(json, false);
}
exports.StatementDetailsStatisticsTypeFromJSON = StatementDetailsStatisticsTypeFromJSON;
function StatementDetailsStatisticsTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'detailsByResort': !(0, runtime_1.exists)(json, 'detailsByResort') ? undefined : (json['detailsByResort'].map(StatementDetailsStatisticType_1.StatementDetailsStatisticTypeFromJSON)),
        'detailsByChannel': !(0, runtime_1.exists)(json, 'detailsByChannel') ? undefined : (json['detailsByChannel'].map(StatementDetailsStatisticType_1.StatementDetailsStatisticTypeFromJSON)),
        'detailsByChannelType': !(0, runtime_1.exists)(json, 'detailsByChannelType') ? undefined : (json['detailsByChannelType'].map(StatementDetailsStatisticType_1.StatementDetailsStatisticTypeFromJSON)),
        'detailsByFeeType': !(0, runtime_1.exists)(json, 'detailsByFeeType') ? undefined : (json['detailsByFeeType'].map(StatementDetailsStatisticType_1.StatementDetailsStatisticTypeFromJSON)),
    };
}
exports.StatementDetailsStatisticsTypeFromJSONTyped = StatementDetailsStatisticsTypeFromJSONTyped;
function StatementDetailsStatisticsTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'detailsByResort': value.detailsByResort === undefined ? undefined : (value.detailsByResort.map(StatementDetailsStatisticType_1.StatementDetailsStatisticTypeToJSON)),
        'detailsByChannel': value.detailsByChannel === undefined ? undefined : (value.detailsByChannel.map(StatementDetailsStatisticType_1.StatementDetailsStatisticTypeToJSON)),
        'detailsByChannelType': value.detailsByChannelType === undefined ? undefined : (value.detailsByChannelType.map(StatementDetailsStatisticType_1.StatementDetailsStatisticTypeToJSON)),
        'detailsByFeeType': value.detailsByFeeType === undefined ? undefined : (value.detailsByFeeType.map(StatementDetailsStatisticType_1.StatementDetailsStatisticTypeToJSON)),
    };
}
exports.StatementDetailsStatisticsTypeToJSON = StatementDetailsStatisticsTypeToJSON;
