/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { RelationshipInfoSummaryType } from './RelationshipInfoSummaryType';
/**
 * Contains a collection of profiles that have a relationship with this profile.
 * @export
 * @interface ProfileTypeRelationshipsSummary
 */
export interface ProfileTypeRelationshipsSummary {
    /**
     * A collection of the profiles summary that have a relationship with this profile.
     * @type {Array<RelationshipInfoSummaryType>}
     * @memberof ProfileTypeRelationshipsSummary
     */
    relationship?: Array<RelationshipInfoSummaryType>;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof ProfileTypeRelationshipsSummary
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof ProfileTypeRelationshipsSummary
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof ProfileTypeRelationshipsSummary
     */
    count?: number;
}
/**
 * Check if a given object implements the ProfileTypeRelationshipsSummary interface.
 */
export declare function instanceOfProfileTypeRelationshipsSummary(value: object): boolean;
export declare function ProfileTypeRelationshipsSummaryFromJSON(json: any): ProfileTypeRelationshipsSummary;
export declare function ProfileTypeRelationshipsSummaryFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileTypeRelationshipsSummary;
export declare function ProfileTypeRelationshipsSummaryToJSON(value?: ProfileTypeRelationshipsSummary | null): any;
