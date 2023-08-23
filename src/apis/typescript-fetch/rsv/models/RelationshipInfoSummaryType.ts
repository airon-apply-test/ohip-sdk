/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { RelationshipProfileSummaryType } from './RelationshipProfileSummaryType';
import {
    RelationshipProfileSummaryTypeFromJSON,
    RelationshipProfileSummaryTypeFromJSONTyped,
    RelationshipProfileSummaryTypeToJSON,
} from './RelationshipProfileSummaryType';

/**
 * RelationshipInfoSummaryType contains information about the associations between and among individuals, companies, travel agents, groups, sources, and contact profiles.
 * @export
 * @interface RelationshipInfoSummaryType
 */
export interface RelationshipInfoSummaryType {
    /**
     * 
     * @type {RelationshipProfileSummaryType}
     * @memberof RelationshipInfoSummaryType
     */
    relationshipProfile?: RelationshipProfileSummaryType;
    /**
     * 
     * @type {object}
     * @memberof RelationshipInfoSummaryType
     */
    masterAccountDetails?: object;
    /**
     * Relationship identifier.
     * @type {string}
     * @memberof RelationshipInfoSummaryType
     */
    relationshipID?: string;
    /**
     * Indicates the type of relationship the current profile(Source Profile) has with the related profile(Target Profile).
     * @type {string}
     * @memberof RelationshipInfoSummaryType
     */
    sourceRelation?: string;
    /**
     * Displays the description of relationship the current profile(Source Profile) has with the related profile(Target Profile).
     * @type {string}
     * @memberof RelationshipInfoSummaryType
     */
    sourceRelationDescription?: string;
    /**
     * Displays the type of relationship the Related profile(Target Profile) has with the current profile(Source Profile).
     * @type {string}
     * @memberof RelationshipInfoSummaryType
     */
    targetRelation?: string;
    /**
     * Displays the description of the target relation(Target Profile).
     * @type {string}
     * @memberof RelationshipInfoSummaryType
     */
    targetRelationDescription?: string;
}

/**
 * Check if a given object implements the RelationshipInfoSummaryType interface.
 */
export function instanceOfRelationshipInfoSummaryType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RelationshipInfoSummaryTypeFromJSON(json: any): RelationshipInfoSummaryType {
    return RelationshipInfoSummaryTypeFromJSONTyped(json, false);
}

export function RelationshipInfoSummaryTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RelationshipInfoSummaryType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'relationshipProfile': !exists(json, 'relationshipProfile') ? undefined : RelationshipProfileSummaryTypeFromJSON(json['relationshipProfile']),
        'masterAccountDetails': !exists(json, 'masterAccountDetails') ? undefined : json['masterAccountDetails'],
        'relationshipID': !exists(json, 'relationshipID') ? undefined : json['relationshipID'],
        'sourceRelation': !exists(json, 'sourceRelation') ? undefined : json['sourceRelation'],
        'sourceRelationDescription': !exists(json, 'sourceRelationDescription') ? undefined : json['sourceRelationDescription'],
        'targetRelation': !exists(json, 'targetRelation') ? undefined : json['targetRelation'],
        'targetRelationDescription': !exists(json, 'targetRelationDescription') ? undefined : json['targetRelationDescription'],
    };
}

export function RelationshipInfoSummaryTypeToJSON(value?: RelationshipInfoSummaryType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'relationshipProfile': RelationshipProfileSummaryTypeToJSON(value.relationshipProfile),
        'masterAccountDetails': value.masterAccountDetails,
        'relationshipID': value.relationshipID,
        'sourceRelation': value.sourceRelation,
        'sourceRelationDescription': value.sourceRelationDescription,
        'targetRelation': value.targetRelation,
        'targetRelationDescription': value.targetRelationDescription,
    };
}

