/**
 * OPERA Cloud Xchange Interface OXI API
 * APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 23.0.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { InstanceLink } from './InstanceLink';
import type { MessageStatisticsReports } from './MessageStatisticsReports';
import type { WarningType } from './WarningType';
/**
 * Response element for Downloaded Messages statistics.
 * @export
 * @interface MessageStatistics
 */
export interface MessageStatistics {
    /**
     *
     * @type {MessageStatisticsReports}
     * @memberof MessageStatistics
     */
    reports?: MessageStatisticsReports;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof MessageStatistics
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof MessageStatistics
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the MessageStatistics interface.
 */
export declare function instanceOfMessageStatistics(value: object): boolean;
export declare function MessageStatisticsFromJSON(json: any): MessageStatistics;
export declare function MessageStatisticsFromJSONTyped(json: any, ignoreDiscriminator: boolean): MessageStatistics;
export declare function MessageStatisticsToJSON(value?: MessageStatistics | null): any;
