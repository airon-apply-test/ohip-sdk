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
 * Represents the site plan section record.
 * @export
 * @interface SitePlanSectionType
 */
export interface SitePlanSectionType {
    /**
     * The site plan section code.
     * @type {string}
     * @memberof SitePlanSectionType
     */
    sectionCode?: string;
    /**
     * The section type.
     * @type {string}
     * @memberof SitePlanSectionType
     */
    sectionType?: string;
    /**
     * The link code for the section.
     * @type {string}
     * @memberof SitePlanSectionType
     */
    linkCode?: string;
    /**
     * The coordinates for this section.
     * @type {string}
     * @memberof SitePlanSectionType
     */
    coordinates?: string;
    /**
     * The description of this section.
     * @type {string}
     * @memberof SitePlanSectionType
     */
    description?: string;
}

/**
 * Check if a given object implements the SitePlanSectionType interface.
 */
export function instanceOfSitePlanSectionType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SitePlanSectionTypeFromJSON(json: any): SitePlanSectionType {
    return SitePlanSectionTypeFromJSONTyped(json, false);
}

export function SitePlanSectionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): SitePlanSectionType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'sectionCode': !exists(json, 'sectionCode') ? undefined : json['sectionCode'],
        'sectionType': !exists(json, 'sectionType') ? undefined : json['sectionType'],
        'linkCode': !exists(json, 'linkCode') ? undefined : json['linkCode'],
        'coordinates': !exists(json, 'coordinates') ? undefined : json['coordinates'],
        'description': !exists(json, 'description') ? undefined : json['description'],
    };
}

export function SitePlanSectionTypeToJSON(value?: SitePlanSectionType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'sectionCode': value.sectionCode,
        'sectionType': value.sectionType,
        'linkCode': value.linkCode,
        'coordinates': value.coordinates,
        'description': value.description,
    };
}

