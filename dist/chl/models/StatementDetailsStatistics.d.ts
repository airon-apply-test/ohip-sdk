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
import type { InstanceLink } from './InstanceLink';
import type { StatementDetailsStatisticsType } from './StatementDetailsStatisticsType';
import type { WarningType } from './WarningType';
/**
 * Response object of the channel account statement details statistics fetch request.
 * @export
 * @interface StatementDetailsStatistics
 */
export interface StatementDetailsStatistics {
    /**
     *
     * @type {StatementDetailsStatisticsType}
     * @memberof StatementDetailsStatistics
     */
    channelStatementDetails?: StatementDetailsStatisticsType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof StatementDetailsStatistics
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof StatementDetailsStatistics
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the StatementDetailsStatistics interface.
 */
export declare function instanceOfStatementDetailsStatistics(value: object): boolean;
export declare function StatementDetailsStatisticsFromJSON(json: any): StatementDetailsStatistics;
export declare function StatementDetailsStatisticsFromJSONTyped(json: any, ignoreDiscriminator: boolean): StatementDetailsStatistics;
export declare function StatementDetailsStatisticsToJSON(value?: StatementDetailsStatistics | null): any;
