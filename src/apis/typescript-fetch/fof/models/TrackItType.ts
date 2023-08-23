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
import type { CodeDescriptionType } from './CodeDescriptionType';
import {
    CodeDescriptionTypeFromJSON,
    CodeDescriptionTypeFromJSONTyped,
    CodeDescriptionTypeToJSON,
} from './CodeDescriptionType';
import type { URLType } from './URLType';
import {
    URLTypeFromJSON,
    URLTypeFromJSONTyped,
    URLTypeToJSON,
} from './URLType';

/**
 * Identifies the kind of Parcel, Baggage, or Lost items or Valet-managed vehicles or services.
 * @export
 * @interface TrackItType
 */
export interface TrackItType {
    /**
     * 
     * @type {CodeDescriptionType}
     * @memberof TrackItType
     */
    type?: CodeDescriptionType;
    /**
     * 
     * @type {URLType}
     * @memberof TrackItType
     */
    url?: URLType;
}

/**
 * Check if a given object implements the TrackItType interface.
 */
export function instanceOfTrackItType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function TrackItTypeFromJSON(json: any): TrackItType {
    return TrackItTypeFromJSONTyped(json, false);
}

export function TrackItTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TrackItType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'type': !exists(json, 'type') ? undefined : CodeDescriptionTypeFromJSON(json['type']),
        'url': !exists(json, 'url') ? undefined : URLTypeFromJSON(json['url']),
    };
}

export function TrackItTypeToJSON(value?: TrackItType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'type': CodeDescriptionTypeToJSON(value.type),
        'url': URLTypeToJSON(value.url),
    };
}

