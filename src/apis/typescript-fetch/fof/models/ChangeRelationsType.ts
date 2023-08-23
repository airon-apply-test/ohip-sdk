/* tslint:disable */
/* eslint-disable */
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

import { exists, mapValues } from '../runtime';
/**
 * Detailed information regarding the changes made to relationships for the profile.
 * @export
 * @interface ChangeRelationsType
 */
export interface ChangeRelationsType {
    /**
     * Relationship identifier.
     * @type {string}
     * @memberof ChangeRelationsType
     */
    id?: string;
    /**
     * Indicates the type of relationship the current profile(Source Profile) has with the related profile(Target Profile).
     * @type {string}
     * @memberof ChangeRelationsType
     */
    relation?: string;
    /**
     * Displays the description of relationship the current profile(Source Profile) has with the related profile(Target Profile).
     * @type {string}
     * @memberof ChangeRelationsType
     */
    relationDescription?: string;
    /**
     * Displays the type of relationship the Related profile(Target Profile) has with the current profile(Source Profile).
     * @type {string}
     * @memberof ChangeRelationsType
     */
    targetRelation?: string;
    /**
     * Displays the description of the target relation(Target Profile).
     * @type {string}
     * @memberof ChangeRelationsType
     */
    targetRelationDescription?: string;
}

/**
 * Check if a given object implements the ChangeRelationsType interface.
 */
export function instanceOfChangeRelationsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ChangeRelationsTypeFromJSON(json: any): ChangeRelationsType {
    return ChangeRelationsTypeFromJSONTyped(json, false);
}

export function ChangeRelationsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChangeRelationsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'relation': !exists(json, 'relation') ? undefined : json['relation'],
        'relationDescription': !exists(json, 'relationDescription') ? undefined : json['relationDescription'],
        'targetRelation': !exists(json, 'targetRelation') ? undefined : json['targetRelation'],
        'targetRelationDescription': !exists(json, 'targetRelationDescription') ? undefined : json['targetRelationDescription'],
    };
}

export function ChangeRelationsTypeToJSON(value?: ChangeRelationsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'relation': value.relation,
        'relationDescription': value.relationDescription,
        'targetRelation': value.targetRelation,
        'targetRelationDescription': value.targetRelationDescription,
    };
}

