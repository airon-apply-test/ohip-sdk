"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.StatisticsReportTypeToJSON = exports.StatisticsReportTypeFromJSONTyped = exports.StatisticsReportTypeFromJSON = exports.StatisticsReportType = void 0;
/**
 * Indicate to get revenue of Stay Records.
 * @export
 */
exports.StatisticsReportType = {
    Reservation: 'Reservation',
    Revenue: 'Revenue',
    ProfileStayRecords: 'ProfileStayRecords',
    DetailStayRecordsRevenue: 'DetailStayRecordsRevenue'
};
function StatisticsReportTypeFromJSON(json) {
    return StatisticsReportTypeFromJSONTyped(json, false);
}
exports.StatisticsReportTypeFromJSON = StatisticsReportTypeFromJSON;
function StatisticsReportTypeFromJSONTyped(json, ignoreDiscriminator) {
    return json;
}
exports.StatisticsReportTypeFromJSONTyped = StatisticsReportTypeFromJSONTyped;
function StatisticsReportTypeToJSON(value) {
    return value;
}
exports.StatisticsReportTypeToJSON = StatisticsReportTypeToJSON;
