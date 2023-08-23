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
/**
 * Post the package on certain nights of the week.
 * @export
 * @interface PackagePostingRhythmTypeCertainNightsOfTheWeek
 */
export interface PackagePostingRhythmTypeCertainNightsOfTheWeek {
    /**
     *
     * @type {boolean}
     * @memberof PackagePostingRhythmTypeCertainNightsOfTheWeek
     */
    sunday?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof PackagePostingRhythmTypeCertainNightsOfTheWeek
     */
    monday?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof PackagePostingRhythmTypeCertainNightsOfTheWeek
     */
    tuesday?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof PackagePostingRhythmTypeCertainNightsOfTheWeek
     */
    wednesday?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof PackagePostingRhythmTypeCertainNightsOfTheWeek
     */
    thursday?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof PackagePostingRhythmTypeCertainNightsOfTheWeek
     */
    friday?: boolean;
    /**
     *
     * @type {boolean}
     * @memberof PackagePostingRhythmTypeCertainNightsOfTheWeek
     */
    saturday?: boolean;
}
/**
 * Check if a given object implements the PackagePostingRhythmTypeCertainNightsOfTheWeek interface.
 */
export declare function instanceOfPackagePostingRhythmTypeCertainNightsOfTheWeek(value: object): boolean;
export declare function PackagePostingRhythmTypeCertainNightsOfTheWeekFromJSON(json: any): PackagePostingRhythmTypeCertainNightsOfTheWeek;
export declare function PackagePostingRhythmTypeCertainNightsOfTheWeekFromJSONTyped(json: any, ignoreDiscriminator: boolean): PackagePostingRhythmTypeCertainNightsOfTheWeek;
export declare function PackagePostingRhythmTypeCertainNightsOfTheWeekToJSON(value?: PackagePostingRhythmTypeCertainNightsOfTheWeek | null): any;
