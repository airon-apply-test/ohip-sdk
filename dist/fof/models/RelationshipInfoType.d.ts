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
import type { ChangeRelationsType } from './ChangeRelationsType';
import type { MasterAccountInfoType } from './MasterAccountInfoType';
import type { RelationshipProfileType } from './RelationshipProfileType';
/**
 * Relationship Type contains information about the associations between and among individuals, companies, travel agents, groups, sources, and contact profiles.
 * @export
 * @interface RelationshipInfoType
 */
export interface RelationshipInfoType {
    /**
     *
     * @type {ChangeRelationsType}
     * @memberof RelationshipInfoType
     */
    changeRelationship?: ChangeRelationsType;
    /**
     *
     * @type {RelationshipProfileType}
     * @memberof RelationshipInfoType
     */
    relationshipProfile?: RelationshipProfileType;
    /**
     *
     * @type {MasterAccountInfoType}
     * @memberof RelationshipInfoType
     */
    masterAccountInfo?: MasterAccountInfoType;
    /**
     * Relationship identifier.
     * @type {string}
     * @memberof RelationshipInfoType
     */
    id?: string;
    /**
     * Indicates the type of relationship the current profile(Source Profile) has with the related profile(Target Profile).
     * @type {string}
     * @memberof RelationshipInfoType
     */
    relation?: string;
    /**
     * Displays the description of relationship the current profile(Source Profile) has with the related profile(Target Profile).
     * @type {string}
     * @memberof RelationshipInfoType
     */
    relationDescription?: string;
    /**
     * Displays the type of relationship the Related profile(Target Profile) has with the current profile(Source Profile).
     * @type {string}
     * @memberof RelationshipInfoType
     */
    targetRelation?: string;
    /**
     * Displays the description of the target relation(Target Profile).
     * @type {string}
     * @memberof RelationshipInfoType
     */
    targetRelationDescription?: string;
}
/**
 * Check if a given object implements the RelationshipInfoType interface.
 */
export declare function instanceOfRelationshipInfoType(value: object): boolean;
export declare function RelationshipInfoTypeFromJSON(json: any): RelationshipInfoType;
export declare function RelationshipInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RelationshipInfoType;
export declare function RelationshipInfoTypeToJSON(value?: RelationshipInfoType | null): any;
