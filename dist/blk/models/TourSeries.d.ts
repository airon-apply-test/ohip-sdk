/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { InstanceLink } from './InstanceLink';
import type { ResponseInstructionType } from './ResponseInstructionType';
import type { TourSeriesType } from './TourSeriesType';
import type { WarningType } from './WarningType';
/**
 * Request for creating a tour series based on an existing block.
 * @export
 * @interface TourSeries
 */
export interface TourSeries {
    /**
     *
     * @type {TourSeriesType}
     * @memberof TourSeries
     */
    criteria?: TourSeriesType;
    /**
     *
     * @type {ResponseInstructionType}
     * @memberof TourSeries
     */
    responseInstruction?: ResponseInstructionType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof TourSeries
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof TourSeries
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the TourSeries interface.
 */
export declare function instanceOfTourSeries(value: object): boolean;
export declare function TourSeriesFromJSON(json: any): TourSeries;
export declare function TourSeriesFromJSONTyped(json: any, ignoreDiscriminator: boolean): TourSeries;
export declare function TourSeriesToJSON(value?: TourSeries | null): any;
